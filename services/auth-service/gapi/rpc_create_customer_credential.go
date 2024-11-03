package gapi

import (
	"context"

	db "github.com/tunvx/simplebank/auth/db/sqlc"
	"github.com/tunvx/simplebank/auth/gapi/val"
	pb "github.com/tunvx/simplebank/grpc/pb/auth"
	cuspb "github.com/tunvx/simplebank/grpc/pb/manage/customer"
	errdb "github.com/tunvx/simplebank/pkg/errs/db"
	errga "github.com/tunvx/simplebank/pkg/errs/gapi"
	"github.com/tunvx/simplebank/pkg/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) CreateCustomerCredential(ctx context.Context, req *pb.CreateCustomerCredentialRequest) (*pb.CreateCustomerCredentialResponse, error) {
	// 1. Validate params
	violations := validateCreateCustomerCredentialRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	// 2. Fetch the customer data for the response
	cusRsp, err := service.customerClient.IGetCustomerByRid(ctx, &cuspb.IGetCustomerByRidRequest{
		CustomerRid: req.CustomerRid,
	})
	if err != nil {
		return nil, err
	}

	// 3. Hash the password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	// 4. Execute the creation of customer credential in the database
	arg := db.CreateCustomerCredentialParams{
		CustomerID:     cusRsp.Customer.CustomerId,
		Username:       req.Username,
		HashedPassword: hashedPassword,
	}
	_, err = service.store.CreateCustomerCredential(ctx, arg)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.UniqueViolation {
			return nil, status.Errorf(codes.NotFound, "customer credential for user ( %s ) already exists in database", req.Username)
		}
		return nil, status.Errorf(codes.Internal, "failed to create customer credential for user ( %s ) into database: %s", req.Username, err)
	}

	// 5. Prepare and return the response
	rsp := &pb.CreateCustomerCredentialResponse{
		IsCreated: true,
	}
	return rsp, nil
}

func validateCreateCustomerCredentialRequest(req *pb.CreateCustomerCredentialRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateCustomerRID(req.GetCustomerRid()); err != nil {
		violations = append(violations, errga.FieldViolation("customer_rid", err))
	}

	// Validate User Name
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, errga.FieldViolation("user_name", err))
	}

	// Validate Password
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, errga.FieldViolation("password", err))
	}

	return violations
}
