package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/tunvx/simplebank/authsrv/db/sqlc"
	"github.com/tunvx/simplebank/authsrv/val"
	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/auth"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) UpdateCustomerCredential(ctx context.Context, req *pb.UpdateCustomerCredentialRequest) (*pb.UpdateCustomerCredentialResponse, error) {
	// 1. Authorize the user
	authPayload, err := service.authorizeUser(ctx, []string{util.CustomerRole})
	if err != nil {
		return nil, err
	}

	// 2. Validate params
	violations := validateUpdateCustomerCredentialRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	customerId := authPayload.UserID

	// 3. Prepare the customer credential parameters
	arg := db.UpdateCustomerCredentialParams{
		CustomerID: customerId,
		ShardID: pgtype.Int4{
			Int32: req.GetShardId(),
			Valid: req.ShardId != nil,
		},
		Username: pgtype.Text{
			String: req.GetUsername(),
			Valid:  req.Username != nil,
		},
	}
	if req.Password != nil {
		// Hash the password
		hashedPassword, err := util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
		}
		arg.HashedPassword = pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		}
	}

	// 4. Execute the update of customer credential in the database
	_, err = service.store.UpdateCustomerCredential(ctx, arg)
	if err != nil {
		if err == errdb.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "customer credential not found for customer id ( %d ) ", authPayload.UserID)
		}
		return nil, status.Errorf(codes.Internal, "failed to update customer credential for customer id ( %d ): %s", authPayload.UserID, err)
	}

	// 5. Prepare and return the response
	response := &pb.UpdateCustomerCredentialResponse{
		IsUpdated: true,
	}
	return response, nil
}

func validateUpdateCustomerCredentialRequest(req *pb.UpdateCustomerCredentialRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Shard ID
	if err := val.ValidateShardID(req.GetShardId()); err != nil {
		violations = append(violations, errga.FieldViolation("shard_id", err))
	}

	// Validate User Name
	if req.Username != nil {
		if err := val.ValidateUsername(req.GetUsername()); err != nil {
			violations = append(violations, errga.FieldViolation("user_name", err))
		}
	}

	// Validate Password
	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, errga.FieldViolation("password", err))
		}
	}

	return violations
}
