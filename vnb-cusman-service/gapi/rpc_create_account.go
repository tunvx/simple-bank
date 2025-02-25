package gapi

import (
	"context"

	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	"github.com/tunvx/simplebank/cusmansrv/val"
	pb "github.com/tunvx/simplebank/grpc/pb/cusman/account"
	"github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	// 1. Validate params and authorize user
	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.CustomerRole})
	if err != nil {
		return nil, err
	}

	// 2. Create new account in original db
	accountID := util.ConvertAccNumberToInt64(req.GetAccountNumber())
	customerId := authPayload.UserID
	shardID := authPayload.ShardID - 1

	_, err = service.shardmanClient.InsertAccountShard(ctx, 
		&shardman.InsertAccountShardRequest{
			AccountId: accountID,
			CustomerId: customerId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create acc shard for account ( %s ): %s", req.GetAccountNumber(), err)
	}
	
	arg := db.CreateAccountParams{
		AccountID:      accountID,
		CustomerID:     authPayload.UserID,
		CurrentBalance: 0,
		CurrencyType:   db.Currencytype(req.CurrencyType),
		AccountStatus:  db.AccountstatusActive,
	}

	account, err := service.stores[shardID].CreateAccount(ctx, arg)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.UniqueViolation {
			return nil, status.Errorf(codes.NotFound, "account ( %s ) already exists in database", req.GetAccountNumber())
		}
		return nil, status.Errorf(codes.Internal, "failed to create account ( %s ) into database: %s", req.GetAccountNumber(), err)
	}

	response := &pb.CreateAccountResponse{
		Account: convertAccount(account),
	}
	return response, nil
}

func validateCreateAccountRequest(req *pb.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
