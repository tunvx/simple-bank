package db

import "context"

type CreateCustomerTxParams struct {
	CreateCustomerParams
	AfterCreate func(customer Customer) error
}

type CreateCustomerTxResult struct {
	Customer Customer
}

func (store *SQLStore) CreateCustomerTx(ctx context.Context, arg CreateCustomerTxParams) (CreateCustomerTxResult, error) {
	var result CreateCustomerTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Customer, err = q.CreateCustomer(ctx, arg.CreateCustomerParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.Customer)
	})

	return result, err
}
