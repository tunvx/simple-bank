package gapi

import (
	"context"

	"github.com/rs/zerolog/log"
	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	pb "github.com/tunvx/simplebank/grpc/pb/moneytransfer"
	"github.com/tunvx/simplebank/grpc/pb/shardman"
	"github.com/tunvx/simplebank/moneytransfersrv/cache"
	"github.com/tunvx/simplebank/moneytransfersrv/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Allow access only to bank customers or other banks
func (service *Service) CheckAccount(ctx context.Context, req *pb.CheckAccountRequest) (*pb.CheckAccountResponse, error) {
	// 1. Validate params and authorize user
	violations := validateCheckAccount(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}
	_, err := service.authorizeUser(ctx,
		[]string{util.CustomerRole, util.BankerRole, util.IServiceRole, util.OtherBanks})
	if err != nil {
		return nil, err
	}

	// 2. Get account info and check it
	accountInfo, err := service.checkAccount(ctx, req.GetAccountNumber(), req.GetCurrencyType())
	if err != nil {
		return nil, err
	}

	// 4. Return response
	response := &pb.CheckAccountResponse{
		IsValid:   true,
		OwnerName: accountInfo.OwnerName,
	}
	return response, nil
}

func (service *Service) checkAccount(
	ctx context.Context,
	accountNumber string,
	currencyTypeRequired string,
) (*cache.AccountInfo, error) {
	accountInfo, err := service.getAccountInfo(ctx, accountNumber)
	if err != nil {
		return nil, err
	}

	if accountInfo.AccountStatus != db.AccountstatusActive {
		return nil, status.Errorf(codes.Internal, "MoneyTransfer Service: Check account failed -> Account ( %s ) is not active", accountNumber)
	}
	if string(accountInfo.CurrencyType) != currencyTypeRequired {
		return nil, status.Errorf(codes.Internal, "MoneyTransfer Service: Check account failed -> Account ( %s ) don't match with currency required ", accountNumber)
	}
	return accountInfo, nil
}

func (service *Service) getAccountInfo(
	ctx context.Context,
	accountNumber string,
) (*cache.AccountInfo, error) {
	var accountInfoPointer *cache.AccountInfo

	// 1. Get account in cache first (ignore errors if any), return if found
	accountId := util.ConvertAccNumberToInt64(accountNumber)

	accountInfoPointer, err := service.cache.GetCacheAccountInfo(ctx, accountId)
	if accountInfoPointer != nil && err == nil {
		return accountInfoPointer, nil
	} // else if err != nil, query from database

	// 2. Cache miss (account not in cache), so, get shard id for account from database
	accShard, err := service.shardmanClient.LookupAccountShard(ctx, &shardman.LookupAccountShardRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "MoneyTransfer Service: Failed to get shard id for account ( %s ): %s", accountNumber, err)
	}

	// 3. Get account based shard_id
	shardId := int(accShard.ShardId) - 1
	account, err := service.stores[shardId].GetAccountForCheck(ctx, accountId)
	if err != nil {
		switch errdb.ErrorCode(err) {
		case errdb.RecordNotFound:
			return nil, status.Errorf(codes.NotFound, "Database: Account ( %s ) not found in shard ( %d )", accountNumber, shardId)
		default:
			return nil, status.Errorf(codes.Internal, "Database: Failed to retrieve account ( %s ) in shard ( %d ): %s", accountNumber, shardId, err)
		}
	}

	accountInfoPointer = convertAccountInfo(account, int(accShard.ShardId))
	log.Debug().Msgf("Database: Successfully retrieved account ( %s ) from shard ( %d )", accountNumber, shardId)

	// 4. Set account to cache (ignore errors if any)
	_ = service.cache.SetCacheAccountInfo(ctx, accountInfoPointer) // ignore errors

	// 5. Return values
	return accountInfoPointer, nil
}

func validateCheckAccount(req *pb.CheckAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
