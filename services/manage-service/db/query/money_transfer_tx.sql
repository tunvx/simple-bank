-- name: CreateMoneyTransferTransaction :one
INSERT INTO money_transfer_transactions (
  amount,
  account_id,
  new_balance,
  description,
  transaction_type,
  transaction_status
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListMoneyTransferTransactionByAccountID :many
SELECT * FROM money_transfer_transactions
WHERE account_id = $1 
ORDER BY transaction_id
LIMIT $2
OFFSET $3;

-- name: GetMoneyTransferTransaction :one
SELECT * FROM money_transfer_transactions
WHERE transaction_id = $1 LIMIT 1;

-- name: UpdateMoneyTransferTransactionStatus :one
UPDATE money_transfer_transactions 
SET 
    transaction_status = $2
WHERE transaction_id = $1
RETURNING *;