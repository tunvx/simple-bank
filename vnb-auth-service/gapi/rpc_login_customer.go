package gapi

import (
	"context"

	db "github.com/tunvx/simplebank/authsrv/db/sqlc"
	"github.com/tunvx/simplebank/authsrv/gapi/val"
	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/auth"
	cuspb "github.com/tunvx/simplebank/grpc/pb/cusman/customer"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *Service) LoginCustomer(ctx context.Context, req *pb.LoginCustomerRequest) (*pb.LoginCustomerResponse, error) {
	// 1. Validate params
	violations := validateLoginCustomerRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	// 2. Retrieve customer credentials & customer management service
	customerCredential, err := service.store.GetCustomerCredential(ctx, req.Username)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "customer credential for user ( %s ) not found in database: ", req.Username)
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve customer credential of user ( %s ) in database: %s", req.Username, err)
	}

	cusRsp, err := service.manageClient.IGetCustomerByID(ctx, &cuspb.IGetCustomerByIDRequest{
		CustomerId: customerCredential.CustomerID,
	})
	if err != nil {
		return nil, err
	}

	// 3. Check that the customer has entered his/her login password correctly
	err = util.CheckPassword(req.Password, customerCredential.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "the password entered is incorrect for user ( %s )", req.Username)
	}

	// 4. Create tokens
	accessToken, accessPayload, err := service.tokenMaker.CreateToken(
		customerCredential.CustomerID,
		util.CustomerRole,
		service.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err)
	}

	refreshToken, refreshPayload, err := service.tokenMaker.CreateToken(
		customerCredential.CustomerID,
		util.BankerRole,
		service.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
	}

	mtdt := service.extractMetadata(ctx)
	session, err := service.store.CreateCustomerSession(ctx, db.CreateCustomerSessionParams{
		SessionID:    refreshPayload.ID,
		CustomerID:   customerCredential.CustomerID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session")
	}

	// 5. Prepare and return the response
	response := &pb.LoginCustomerResponse{
		Customer:              cusRsp.GetCustomer(),
		SessionId:             session.SessionID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return response, nil
}

func validateLoginCustomerRequest(req *pb.LoginCustomerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
