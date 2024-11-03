package gapi

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	cuspb "github.com/tunvx/simplebank/grpc/pb/manage/customer"
	db "github.com/tunvx/simplebank/manage/db/sqlc"
	"github.com/tunvx/simplebank/manage/gapi/val"
	"github.com/tunvx/simplebank/notification/redis"
	errdb "github.com/tunvx/simplebank/pkg/errs/db"
	errga "github.com/tunvx/simplebank/pkg/errs/gapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) CreateCustomer(ctx context.Context, req *cuspb.CreateCustomerRequest) (*cuspb.CreateCustomerResponse, error) {
	// 1. Validate params
	violations := validateCreateCustomerRequest(req)
	if violations != nil {
		// Return error if there are any violations in the request
		return nil, errga.InvalidArgumentError(violations)
	}

	// 2. Parse string dateOfBirth from request to time.Time, assuming it's already validated
	dateOfBirth, _ := time.Parse("2006/01/02", req.GetDateOfBirth())

	// 3. Prepare the arguments for creating a customer in the database
	arg := db.CreateCustomerTxParams{
		CreateCustomerParams: db.CreateCustomerParams{
			CustomerRid:     req.GetCustomerRid(),
			Fullname:        req.GetFullname(),
			DateOfBirth:     dateOfBirth,
			Address:         req.GetAddress(),
			PhoneNumber:     req.GetPhoneNumber(),
			Email:           req.GetEmail(),
			CustomerTier:    db.Customertier(req.GetCustomerTier()),
			CustomerSegment: db.Customersegment(req.GetCustomerSegment()),
			FinancialStatus: db.Financialstatus(req.GetFinancialStatus()),
		},

		// Important step: Distribute/assign the task of sending verification mail for background worker.
		AfterCreate: func(custom db.Customer) error {
			taskPayload := &redis.PayloadSendVerifyEmail{
				CustomerRid: custom.CustomerRid,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(redis.QueueCritical),
			}

			// Time to distribute/assign the task is very small
			return service.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	// 4. Call the database layer to create the customer
	txResult, err := service.store.CreateCustomerTx(ctx, arg)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create user ( %s ): %s", req.GetCustomerRid(), err)
	}

	rsp := &cuspb.CreateCustomerResponse{
		Customer: convertCustomer(txResult.Customer),
	}
	return rsp, nil
}

func validateCreateCustomerRequest(req *cuspb.CreateCustomerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateCustomerRID(req.GetCustomerRid()); err != nil {
		violations = append(violations, errga.FieldViolation("customer_rid", err))
	}

	// Validate Full Name
	if err := val.ValidateFullName(req.GetFullname()); err != nil {
		violations = append(violations, errga.FieldViolation("full_name", err))
	}

	// Validate Date of Birth
	if err := val.ValidateDateOfBirth(req.GetDateOfBirth()); err != nil {
		violations = append(violations, errga.FieldViolation("date_of_birth", err))
	}

	// Validate Address (optional: depending on requirements)
	if err := val.ValidateString(req.GetAddress(), 5, 255); err != nil {
		violations = append(violations, errga.FieldViolation("address", err))
	}

	// Validate Phone Number (ensure it is a valid format and length)
	if err := val.ValidatePhoneNumber(req.GetPhoneNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("phone_number", err))
	}

	// Validate Email
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, errga.FieldViolation("email", err))
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
