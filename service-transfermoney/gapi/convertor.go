package gapi

import (
	accpb "github.com/tunvx/simplebank/grpc/pb/management/account"
	tranpb "github.com/tunvx/simplebank/grpc/pb/transfermoney"
	db "github.com/tunvx/simplebank/management/db/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAccount(account db.Account) *accpb.Account {
	return &accpb.Account{
		AccountNumber:  account.AccountNumber,
		CurrentBalance: account.CurrentBalance,
		CurrencyType:   string(account.CurrencyType),
		CreatedAt:      timestamppb.New(account.CreatedAt),
		Description:    account.Description,
		AccountStatus:  string(account.AccountStatus),
	}
}

func convertMoneyTransferTx(transaction db.MoneyTransferTransaction) *tranpb.MoneyTransferTx {
	return &tranpb.MoneyTransferTx{
		TransactionId:     transaction.TransactionID,
		Amount:            transaction.Amount,
		NewBalance:        transaction.NewBalance,
		AccountId:         transaction.AccountID,
		TransactionTime:   timestamppb.New(transaction.TransactionTime),
		Description:       transaction.Description,
		TransactionType:   string(transaction.TransactionType),
		TransactionStatus: string(transaction.TransactionStatus),
	}
}
