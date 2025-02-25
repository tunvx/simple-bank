package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type VerifyEmailTxParams struct {
	EmailId    uuid.UUID // [16] byte
	SecretCode string
}

type VerifyEmailTxResult struct {
	Customer    Customer
	VerifyEmail VerifyEmail
}

func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		result.Customer, err = q.UpdateCustomer(ctx, UpdateCustomerParams{
			CustomerID: result.VerifyEmail.CustomerID,
			IsEmailVerified: pgtype.Bool{
				Bool:  true,
				Valid: true,
			},
		})
		return err
	})

	return result, err
}
