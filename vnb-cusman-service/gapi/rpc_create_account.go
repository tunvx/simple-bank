package gapi

import (
	"context"

	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	"github.com/tunvx/simplebank/cusmansrv/gapi/val"
	accpb "github.com/tunvx/simplebank/grpc/pb/cusman/account"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) CreateAccount(ctx context.Context, req *accpb.CreateAccountRequest) (*accpb.CreateAccountResponse, error) {
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.CustomerRole})
	if err != nil {
		return nil, err
	}

	userID := authPayload.UserID
	shardID := util.ExtractShardID(userID)

	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	arg := db.CreateAccountParams{
		AccountNumber:  req.AccountNumber,
		CustomerID:     authPayload.UserID,
		CurrentBalance: 0,
		CurrencyType:   db.Currencytype(req.CurrencyType),
		AccountStatus:  db.AccountstatusActive,
	}

	account, err := service.stores[shardID].CreateAccount(ctx, arg)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.UniqueViolation {
			return nil, status.Errorf(codes.NotFound, "account ( %s ) already exists in database", req.AccountNumber)
		}
		return nil, status.Errorf(codes.Internal, "failed to create account ( %s ) into database: %s", req.AccountNumber, err)
	}

	response := &accpb.CreateAccountResponse{
		Account: convertAccount(account),
	}
	return response, nil
}

func validateCreateAccountRequest(req *accpb.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateAccountNumber(req.GetAccountNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("account_number", err))
	}

	// Validate User Name
	if err := val.ValidateCurrency(req.GetCurrencyType()); err != nil {
		violations = append(violations, errga.FieldViolation("currency_type", err))
	}

	return violations
}
