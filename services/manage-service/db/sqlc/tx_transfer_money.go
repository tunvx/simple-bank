package db

import (
	"context"
	"fmt"

	errdb "github.com/tunvx/simplebank/pkg/errs/db"
)

// The `SenderDescription` and `RecipientDescription` fields are used to store
// descriptive information related to the transaction, such as the bank ID,
// account number, and account owner's name. By including these details within
// the description fields, we simplify the structure and facilitate easier
// creation, updating, retrieval, and deletion (CRUD) operations within the database.
// This approach avoids the need for separate columns to store this information,
// making the schema more streamlined.

// CreateCompleteTxParams contains the parameters required to create a complete transaction, including both sending and receiving sides.
type CreateCompleteTxParams struct {
	Amount             int64  `json:"amount"`
	SenderAccountID    int64  `json:"sender_account_id"`
	RecipientAccountID int64  `json:"recipient_account_id"`
	Message            string `json:"message"`
	AfterTransfer      func(amount int64, serder_account Account, recipient_account Account) error
}

type CreateCompleteTxResult struct {
	SenderAccount        Account                  `json:"sender_account"`
	RecipientAccount     Account                  `json:"recipient_account"`
	SendingTransaction   MoneyTransferTransaction `json:"sending_transaction"`
	ReceivingTransaction MoneyTransferTransaction `json:"receiving_transaction"`
}

// CreateCompleteTransaction performs a complete transaction between two accounts.
func (store *SQLStore) CreateCompleteTx(
	ctx context.Context,
	arg CreateCompleteTxParams,
) (CreateCompleteTxResult, error) {
	var result CreateCompleteTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Lock accounts in consistent order to avoid deadlocks
		// Ensure accounts are always locked in ascending order of AccountID
		if arg.SenderAccountID < arg.RecipientAccountID {
			// Case 1: Update amount: Deduct from sender_account at first, add to recipient_account at second
			result.SenderAccount, result.RecipientAccount, err = transferMoneyBetweenTwoAccount(
				ctx, q, arg.SenderAccountID, -arg.Amount, arg.RecipientAccountID, arg.Amount)
		} else {
			// Case 2: Update amount: Add to recipient_account first, deduct from sender_account second
			result.RecipientAccount, result.SenderAccount, err = transferMoneyBetweenTwoAccount(
				ctx, q, arg.RecipientAccountID, arg.Amount, arg.SenderAccountID, -arg.Amount)
		}
		if err != nil {
			return err
		}

		err = validateSenderAccountBalance(result.SenderAccount, arg.Amount)
		if err != nil {
			return err
		}

		// Create sending transaction
		result.SendingTransaction, err = q.CreateMoneyTransferTransaction(ctx,
			CreateMoneyTransferTransactionParams{
				Amount:            -arg.Amount,
				AccountID:         arg.SenderAccountID,
				NewBalance:        result.SenderAccount.CurrentBalance,
				Description:       arg.Message,
				TransactionType:   TransactiontypeSendMoneyInternal,
				TransactionStatus: TransactionstatusCompleted,
			})
		if err != nil {
			return err
		}

		// Create receiving transaction
		result.ReceivingTransaction, err = q.CreateMoneyTransferTransaction(ctx,
			CreateMoneyTransferTransactionParams{
				Amount:            arg.Amount,
				AccountID:         arg.RecipientAccountID,
				NewBalance:        result.RecipientAccount.CurrentBalance,
				Description:       arg.Message,
				TransactionType:   TransactiontypeReceiveMoneyInternal,
				TransactionStatus: TransactionstatusCompleted,
			})
		if err != nil {
			return err
		}

		return arg.AfterTransfer(arg.Amount, result.SenderAccount, result.RecipientAccount)
	})
	return result, err
}

func transferMoneyIO(
	ctx context.Context,
	q *Queries,
	accountID int64,
	amount int64,
) (account Account, err error) {
	account, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		AccountID: accountID,
		Amount:    amount,
	})
	if err != nil {
		return
	}
	return
}

func transferMoneyBetweenTwoAccount(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		AccountID: accountID1,
		Amount:    amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		AccountID: accountID2,
		Amount:    amount2,
	})
	return
}

func validateSenderAccountBalance(senderAccount Account, amount int64) error {
	if senderAccount.CurrentBalance < 0 {
		err := fmt.Errorf("%w: account_number( %v ), current_balance ( %d ), transfer_amount ( %d )",
			errdb.ErrInsufficientFunds, senderAccount.AccountNumber, senderAccount.CurrentBalance+amount, amount)
		return err
	}
	return nil
}
