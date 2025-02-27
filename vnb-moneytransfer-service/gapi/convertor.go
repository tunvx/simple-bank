package gapi

import (
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	accpb "github.com/tunvx/simplebank/grpc/pb/cusman/account"
	tranpb "github.com/tunvx/simplebank/grpc/pb/moneytransfer"
	"github.com/tunvx/simplebank/moneytransfersrv/cache"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAccountInfo(account db.GetAccountForCheckRow, shardID int) *cache.AccountInfo {
	return &cache.AccountInfo{
		AccountID:     account.AccountID,
		ShardId:       shardID,
		CustomerID:    account.CustomerID,
		OwnerName:     account.OwnerName,
		CurrencyType:  account.CurrencyType,
		AccountStatus: account.AccountStatus,
	}
}

func convertAccount(account db.Account) *accpb.Account {
	return &accpb.Account{
		AccountId:      account.AccountID,
		CurrentBalance: account.CurrentBalance,
		CurrencyType:   string(account.CurrencyType),
		CreatedAt:      timestamppb.New(account.CreatedAt),
		Description:    account.Description,
		AccountStatus:  string(account.AccountStatus),
	}
}

func convertAccountTransaction(transaction db.AccountTransaction) *tranpb.AccountTransaction {
	transactionIdStr, _ := util.ConvertUUIDToString(transaction.TransactionID)
	return &tranpb.AccountTransaction{
		TransactionId:     transactionIdStr,
		Amount:            transaction.Amount,
		NewBalance:        transaction.NewBalance,
		AccountId:         transaction.AccountID,
		TransactionTime:   timestamppb.New(transaction.TransactionTime),
		Description:       transaction.Description,
		TransactionType:   string(transaction.TransactionType),
		TransactionStatus: string(transaction.TransactionStatus),
	}
}
