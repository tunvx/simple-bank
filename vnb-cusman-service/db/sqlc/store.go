package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier

	CreateCustomerTx(ctx context.Context, arg CreateCustomerTxParams) (CreateCustomerTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)

	// CreateCompleteTransaction creates a complete transaction by deducting money from account A
	// and crediting money to account B. This operation is final and no further actions are required.
	CreateCompleteTx(ctx context.Context, arg CreateCompleteTxParams) (CreateCompleteTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// Close closes the database connection pool
func (s *SQLStore) Close() {
	s.connPool.Close()
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
