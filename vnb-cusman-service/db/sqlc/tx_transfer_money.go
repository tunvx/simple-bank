package db

import (
	"context"

	"github.com/google/uuid"
)

type CreateTxTransferMoneyParams struct {
	TransactionID	uuid.UUID `json:"transaction_id"`
	Amount        	int64  `json:"amount"`
	SourceAccID 	int64  `json:"source_acc_id"`
	Description 	string `json:"description"`
	TransactionType Transactiontype `json:"transaction_type"`
}

type CreateTxTransferMoneyResult struct {
	SourceAccount   Account            `json:"source_account"`
	SendingTransaction AccountTransaction `json:"sending_transaction"`
}

// CreateTxTransferMoney performs a transaction transfer money.
func (store *SQLStore) CreateTxTransferMoney(
	ctx context.Context,
	arg CreateTxTransferMoneyParams,
) (CreateTxTransferMoneyResult, error) {
	var resutl CreateTxTransferMoneyResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		resutl.SourceAccount, err = transferMoneyIO(ctx, q, arg.SourceAccID, -arg.Amount)
		if err != nil {
			return err
		}

		resutl.SendingTransaction, err = q.CreateAccountTransaction(ctx, 
			CreateAccountTransactionParams{
				TransactionID: arg.TransactionID,
				Amount: -arg.Amount,
				NewBalance: resutl.SourceAccount.CurrentBalance,
				Description: arg.Description,
				TransactionType: arg.TransactionType,
				TransactionStatus: TransactionstatusCompleted,
			})
			
		return err
	})
	return resutl, err
}

// CreateCompleteTxInShardParams contains the parameters required to 
// create a complete transaction, including both sending and receiving sides.
type CreateCompleteTxInShardParams struct {
	Amount             int64  `json:"amount"`
	SourceAccID 	int64  `json:"source_acc_id"`
	BeneficiaryAccID int64  `json:"beneficiary_acc_id"`
	Description 	string `json:"description"`
}

type CreateCompleteTxInShardResult struct {
	SourceAccount        Account                  `json:"source_account"`
	BeneficiaryAccount    Account                  `json:"Beneficiary_account"`
	SendingTransaction   AccountTransaction `json:"sending_transaction"`
	ReceivingTransaction AccountTransaction `json:"receiving_transaction"`
}

// CreateCompleteTxInShard performs a complete transaction between two accounts.
func (store *SQLStore) CreateCompleteTxInShard(
	ctx context.Context,
	arg CreateCompleteTxInShardParams,
) (CreateCompleteTxInShardResult, error) {
	var result CreateCompleteTxInShardResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Lock accounts in consistent order to avoid deadlocks
		// Ensure accounts are always locked in ascending order of AccountID
		if arg.SourceAccID < arg.BeneficiaryAccID {
			// Case 1: Update amount: Deduct from sender_account at first, add to recipient_account at second
			result.SourceAccount, result.BeneficiaryAccount, err = transferMoneyBetweenTwoAccount(
				ctx, q, arg.SourceAccID, -arg.Amount, arg.BeneficiaryAccID, arg.Amount)
		} else {
			// Case 2: Update amount: Add to recipient_account first, deduct from sender_account second
			result.BeneficiaryAccount, result.SourceAccount, err = transferMoneyBetweenTwoAccount(
				ctx, q, arg.BeneficiaryAccID, arg.Amount, arg.SourceAccID, -arg.Amount)
		}
		if err != nil {
			return err
		}

		err = validateSourceAccountBalance(result.SourceAccount, arg.Amount)
		if err != nil {
			return err
		}

		// Create sending transaction
		result.SendingTransaction, err = q.CreateAccountTransaction(ctx,
			CreateAccountTransactionParams{
				Amount:            -arg.Amount,
				AccountID:         arg.SourceAccID,
				NewBalance:        result.SourceAccount.CurrentBalance,
				Description:       arg.Description,
				TransactionType:   TransactiontypeInternalSend,
				TransactionStatus: TransactionstatusCompleted,
			})
		if err != nil {
			return err
		}

		// Create receiving transaction
		result.ReceivingTransaction, err = q.CreateAccountTransaction(ctx,
			CreateAccountTransactionParams{
				Amount:            arg.Amount,
				AccountID:         arg.BeneficiaryAccID,
				NewBalance:        result.BeneficiaryAccount.CurrentBalance,
				Description:       arg.Description,
				TransactionType:   TransactiontypeInternalReceive,
				TransactionStatus: TransactionstatusCompleted,
			})
		if err != nil {
			return err
		}
		return err
	})
	return result, err
}




