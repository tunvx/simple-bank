package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier

	// Management transactions
	CreateCustomerTx(ctx context.Context, arg CreateCustomerTxParams) (CreateCustomerTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)

	// Money transfer transactions
	CreateTxReceiveMoney(ctx context.Context, arg CreateTxReceiveMoneyParams) (CreateTxReceiveMoneyResult, error)
	CreateTxTransferMoney(ctx context.Context, arg CreateTxTransferMoneyParams) (CreateTxTransferMoneyResult, error)
	CreateCompleteTxInShard(ctx context.Context, arg CreateCompleteTxInShardParams) (CreateCompleteTxInShardResult, error)
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
