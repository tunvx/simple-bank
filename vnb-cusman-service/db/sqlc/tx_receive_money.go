package db

import (
	"context"

	"github.com/google/uuid"
)

type CreateTxReceiveMoneyParams struct {
	TransactionID    uuid.UUID `json:"transaction_id"`
	Amount           int64  `json:"amount"`
	BeneficiaryAccID int64  `json:"beneficiary_acc_id"`
	Description string `json:"description"`
	TransactionType Transactiontype `json:"transaction_type"`
}

type CreateTxReceiveMoneyResult struct {
	BeneficiaryAccount   Account            `json:"beneficiary_account"`
	ReceivingTransaction AccountTransaction `json:"receiving_transaction"`
}

// CreateTxReceiveMoney performs a transaction receive money.
func (store *SQLStore) CreateTxReceiveMoney(
	ctx context.Context,
	arg CreateTxReceiveMoneyParams,
) (CreateTxReceiveMoneyResult, error) {
	var resutl CreateTxReceiveMoneyResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		resutl.BeneficiaryAccount, err = transferMoneyIO(ctx, q, arg.BeneficiaryAccID, arg.Amount)
		if err != nil {
			return err
		}

		resutl.ReceivingTransaction, err = q.CreateAccountTransaction(ctx, 
			CreateAccountTransactionParams{
				TransactionID: arg.TransactionID,
				Amount: arg.Amount,
				NewBalance: resutl.BeneficiaryAccount.CurrentBalance,
				Description: arg.Description,
				TransactionType: arg.TransactionType,
				TransactionStatus: TransactionstatusCompleted,
			})
			
		return err
	})
	return resutl, err
}
