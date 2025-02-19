package gapi

import (
	"context"

	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	cuspb "github.com/tunvx/simplebank/grpc/pb/management/customer"
	db "github.com/tunvx/simplebank/management/db/sqlc"
	"github.com/tunvx/simplebank/management/gapi/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) VerifyEmail(ctx context.Context, req *cuspb.VerifyEmailRequest) (*cuspb.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	shardID := util.ExtractShardID(req.GetEmailId())

	txResult, err := service.stores[shardID].VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailId:    req.GetEmailId(),
		SecretCode: req.GetSecretCode(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify email ( %d ): %v", req.GetEmailId(), err)
	}

	rsp := &cuspb.VerifyEmailResponse{
		IsVerified: txResult.Customer.IsEmailVerified,
	}
	return rsp, nil
}

func validateVerifyEmailRequest(req *cuspb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmailId(req.GetEmailId()); err != nil {
		violations = append(violations, errga.FieldViolation("email_id", err))
	}

	if err := val.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, errga.FieldViolation("secret_code", err))
	}

	return violations
}
