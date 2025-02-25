package gapi

import (
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	accpb "github.com/tunvx/simplebank/grpc/pb/cusman/account"
	cuspb "github.com/tunvx/simplebank/grpc/pb/cusman/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
)



func convertCustomer(customer db.Customer) *cuspb.Customer {
	return &cuspb.Customer{
		CustomerRid:     customer.CustomerRid,
		FullName:        customer.FullName,
		DateOfBirth:     customer.DateOfBirth.String(),
		PermanentAddress:         customer.PermanentAddress,
		PhoneNumber:     customer.PhoneNumber,
		EmailAddress:           customer.EmailAddress,
		CustomerTier:    string(customer.CustomerTier),
		CustomerSegment: string(customer.CustomerSegment),
		FinancialStatus: string(customer.FinancialStatus),
	}
}

func convertAccount(account db.Account) *accpb.Account {
	return &accpb.Account{
		AccountId:  account.AccountID,
		CurrentBalance: account.CurrentBalance,
		CurrencyType:   string(account.CurrencyType),
		CreatedAt:      timestamppb.New(account.CreatedAt),
		Description:    account.Description,
		AccountStatus:  string(account.AccountStatus),
	}
}
