package db

import (
	"context"
	"fmt"

	errdb "github.com/tunvx/simplebank/common/errs/db"
)

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

func validateSourceAccountBalance(sourceAccount Account, amount int64) error {
	if sourceAccount.CurrentBalance < 0 {
		err := fmt.Errorf("%w: account_number( %d ), current_balance ( %d ), transfer_amount ( %d )",
			errdb.ErrInsufficientFunds, sourceAccount.AccountID, sourceAccount.CurrentBalance+amount, amount)
		return err
	}
	return nil
}
