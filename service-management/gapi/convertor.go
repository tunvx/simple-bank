package gapi

import (
	accpb "github.com/tunvx/simplebank/grpc/pb/management/account"
	cuspb "github.com/tunvx/simplebank/grpc/pb/management/customer"
	db "github.com/tunvx/simplebank/management/db/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertICustomer(customer db.Customer) *cuspb.ICustomer {
	return &cuspb.ICustomer{
		CustomerId:      customer.CustomerID,
		CustomerRid:     customer.CustomerRid,
		Fullname:        customer.Fullname,
		DateOfBirth:     customer.DateOfBirth.String(),
		Address:         customer.Address,
		PhoneNumber:     customer.PhoneNumber,
		Email:           customer.Email,
		CustomerTier:    string(customer.CustomerTier),
		CustomerSegment: string(customer.CustomerSegment),
		FinancialStatus: string(customer.FinancialStatus),
	}
}

func convertCustomer(customer db.Customer) *cuspb.Customer {
	return &cuspb.Customer{
		CustomerRid:     customer.CustomerRid,
		Fullname:        customer.Fullname,
		DateOfBirth:     customer.DateOfBirth.String(),
		Address:         customer.Address,
		PhoneNumber:     customer.PhoneNumber,
		Email:           customer.Email,
		CustomerTier:    string(customer.CustomerTier),
		CustomerSegment: string(customer.CustomerSegment),
		FinancialStatus: string(customer.FinancialStatus),
	}
}

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
