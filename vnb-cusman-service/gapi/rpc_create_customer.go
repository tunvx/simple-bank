package gapi

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	"github.com/tunvx/simplebank/cusmansrv/val"
	worker "github.com/tunvx/simplebank/cusmansrv/worker/redis"
	pb "github.com/tunvx/simplebank/grpc/pb/cusman/customer"
	"github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	// 1. Validate params
	violations := validateCreateCustomerRequest(req)
	if violations != nil {
		// Return error if there are any violations in the request
		return nil, errga.InvalidArgumentError(violations)
	}

	// 2. Parse string dateOfBirth from request to time.Time, assuming it's already validated
	dateOfBirth, _ := time.Parse("2006/01/02", req.GetDateOfBirth())

	// 3. Create new customer in original db
	newCusShard, err := service.shardmanClient.InsertCustomerShard(ctx, &shardman.InsertCustomerShardRequest{
		CustomerRid: req.GetCustomerRid(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create id for customer ( %s ): %s", req.GetCustomerRid(), err)
	}

	// 4. Prepare the arguments for creating a customer in the database
	arg := db.CreateCustomerTxParams{
		CreateCustomerParams: db.CreateCustomerParams{
			CustomerID:       newCusShard.GetCustomerId(),
			CustomerRid:      req.GetCustomerRid(),
			FullName:         req.GetFullName(),
			DateOfBirth:      dateOfBirth,
			PermanentAddress: req.GetPermanentAddress(),
			PhoneNumber:      req.GetPhoneNumber(),
			EmailAddress:     req.GetEmailAddress(),
			CustomerTier:     db.Customertier(req.GetCustomerTier()),
			CustomerSegment:  db.Customersegment(req.GetCustomerSegment()),
			FinancialStatus:  db.Financialstatus(req.GetFinancialStatus()),
		},

		// Important step: Distribute/assign the task of sending verification mail for background worker.
		AfterCreate: func(custom db.Customer) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				CustomerID: custom.CustomerID,
				ShardID:    newCusShard.ShardId,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}

			// Time to distribute/assign the task is very small
			return service.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	// 5. Call the database layer to create the customer
	shardId := newCusShard.GetShardId() - 1
	txResult, err := service.stores[shardId].CreateCustomerTx(ctx, arg)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "%s", err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create user ( %s ): %s", req.GetCustomerRid(), err)
	}

	// 6. Return Reponse
	rsp := &pb.CreateCustomerResponse{
		Customer: convertCustomer(txResult.Customer),
	}
	return rsp, nil
}

func validateCreateCustomerRequest(req *pb.CreateCustomerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateCustomerRID(req.GetCustomerRid()); err != nil {
		violations = append(violations, errga.FieldViolation("customer_rid", err))
	}

	// Validate Full Name
	if err := val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, errga.FieldViolation("full_name", err))
	}

	// Validate Date of Birth
	if err := val.ValidateDateOfBirth(req.GetDateOfBirth()); err != nil {
		violations = append(violations, errga.FieldViolation("date_of_birth", err))
	}

	// Validate Address (optional: depending on requirements)
	if err := val.ValidateString(req.GetPermanentAddress(), 5, 255); err != nil {
		violations = append(violations, errga.FieldViolation("permanent_address", err))
	}

	// Validate Phone Number (ensure it is a valid format and length)
	if err := val.ValidatePhoneNumber(req.GetPhoneNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("phone_number", err))
	}

	// Validate Email
	if err := val.ValidateEmail(req.GetEmailAddress()); err != nil {
		violations = append(violations, errga.FieldViolation("email_address", err))
	}

	// Validate Customer Tier
	if err := val.ValidateCustomerTier(req.GetCustomerTier()); err != nil {
		violations = append(violations, errga.FieldViolation("customer_tier", err))
	}

	// Validate Customer Segment
	if err := val.ValidateCustomerSegment(req.GetCustomerSegment()); err != nil {
		violations = append(violations, errga.FieldViolation("customer_segment", err))
	}

	// Validate Financial Status
	if err := val.ValidateFinancialStatus(req.GetFinancialStatus()); err != nil {
		violations = append(violations, errga.FieldViolation("financial_status", err))
	}

	return violations
}
