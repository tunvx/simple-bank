package gapi

import (
	"context"

	"github.com/rs/zerolog/log"
	errdb "github.com/tunvx/simplebank/common/errs/db"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/transfermoney"
	db "github.com/tunvx/simplebank/management/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) CheckAccount(ctx context.Context, req *pb.CheckAccountRequest) (*pb.CheckAccountResponse, error) {
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.CustomerRole})
	if err != nil {
		return nil, err
	}

	userID := authPayload.UserID
	shardID := util.ExtractShardID(userID)

	account, err := service.getAndValidateAccount(ctx, shardID, req.AccNumber, req.CurrencyType)
	if err != nil {
		return nil, err
	}

	response := &pb.CheckAccountResponse{
		Account: convertAccount(account),
	}
	return response, nil
}

func (service *Service) getAndValidateAccount(ctx context.Context, shardID int, accountNumber string, currencyType string) (db.Account, error) {
	// Check account in cache first
	account, err := service.cache.GetCacheAccount(ctx, accountNumber)
	if err != nil {
		log.Warn().Msgf("failed to get account (%s) from cache: %s", accountNumber, err)
	}

	if account.AccountNumber != "" {
		log.Info().Msgf("cache hit: account ( %s ) in cache", accountNumber)
	} else {
		log.Info().Msgf("cache miss: account ( %s ) not in cache", accountNumber)

		// Cache miss (account not in cache), so, get from database
		if account.AccountNumber == "" {
			account, err = service.stores[shardID].GetAccountByAccNumber(ctx, accountNumber)
			if err != nil {
				if errdb.ErrorCode(err) == errdb.RecordNotFound {
					return account, status.Errorf(codes.NotFound, "account ( %s ) not found in database", accountNumber)
				}
				return account, status.Errorf(codes.Internal, "failed to retrieve account ( %s ) in database: %s", accountNumber, err)
			}

			// After getting from the database, set account to cache
			err = service.cache.SetCacheAccount(ctx, account)
			if err != nil {
				log.Warn().Msgf("failed to set account %s in cache: %s", accountNumber, err)
			}
		}
	}

	// Check account status
	if account.AccountStatus != db.AccountstatusActive {
		return account, status.Errorf(codes.FailedPrecondition, "account ( %s ) is not active", account.AccountNumber)
	}

	// Check if the account's currency matches the expected currency
	if string(account.CurrencyType) != currencyType {
		return account, status.Errorf(codes.InvalidArgument, "account ( %s ) has a different currency than the transfer request", account.AccountNumber)
	}

	// Here has token (encode) and return it
	// + has int64 customer.CustomerID
	// + has int64 account.AccountID
	// + has bool AllowTransfer
	// + has time ExpiredAt

	return account, nil
}
