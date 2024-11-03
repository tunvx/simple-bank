package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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

func (service *Service) UpdateCustomerCredential(ctx context.Context, req *pb.UpdateCustomerCredentialRequest) (*pb.UpdateCustomerCredentialResponse, error) {
	// 1. Authorize the user
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.CustomerRole})
	if err != nil {
		return nil, err
	}

	// 2. Validate params
	violations := validateUpdateCustomerCredentialRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	// 3. Fetch the customer data for the response
	cusRsp, err := service.customerClient.IGetCustomerByRid(ctx, &cuspb.IGetCustomerByRidRequest{
		CustomerRid: req.CustomerRid,
	})
	if err != nil {
		return nil, err
	}

	// 4. Ensure the authenticated user can only update their own info
	if authPayload.Role == util.CustomerRole && authPayload.UserID != cusRsp.Customer.CustomerId {
		return nil, status.Errorf(codes.PermissionDenied, "user_id cannot match with owner_id")
	}

	// 5. Prepare the customer credential parameters
	arg := db.UpdateCustomerCredentialParams{
		CustomerID: cusRsp.Customer.CustomerId,
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

	// 6. Execute the update of customer credential in the database
	_, err = service.store.UpdateCustomerCredential(ctx, arg)
	if err != nil {
		if err == errdb.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "customer credential not found for customer rid ( %s ) ", req.CustomerRid)
		}
		return nil, status.Errorf(codes.Internal, "failed to update customer credential for customer rid ( %s ): %s", req.CustomerRid, err)
	}

	// 7. Prepare and return the response
	response := &pb.UpdateCustomerCredentialResponse{
		IsUpdated: true,
	}
	return response, nil
}

func validateUpdateCustomerCredentialRequest(req *pb.UpdateCustomerCredentialRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateCustomerRID(req.GetCustomerRid()); err != nil {
		violations = append(violations, errga.FieldViolation("customer_rid", err))
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
