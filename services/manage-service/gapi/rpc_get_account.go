package gapi

import (
	"context"

	accpb "github.com/tunvx/simplebank/grpc/pb/manage/account"
	"github.com/tunvx/simplebank/manage/gapi/val"
	errdb "github.com/tunvx/simplebank/pkg/errs/db"
	errga "github.com/tunvx/simplebank/pkg/errs/gapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) GetAccountByAccNumber(ctx context.Context, req *accpb.GetAccountByAccNumberRequest) (*accpb.GetAccountByAccNumberResponse, error) {
	violations := validateGetAccountByAccNumber(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	account, err := service.store.GetAccountByAccNumber(ctx, req.AccNumber)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "account ( %s ) not found in database", req.AccNumber)
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve account ( %s ) in database: %s", req.AccNumber, err)
	}

	response := &accpb.GetAccountByAccNumberResponse{
		Account: convertAccount(account),
	}
	return response, nil
}

func validateGetAccountByAccNumber(req *accpb.GetAccountByAccNumberRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateAccountNumber(req.GetAccNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("account_number", err))
	}

	return violations
}
