-- name: CreateAccountTransaction :one
INSERT INTO account_transactions (
  transaction_id,
  amount,
  account_id,
  new_balance,
  description,
  transaction_type,
  transaction_status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetAccountTransaction :one
SELECT * FROM account_transactions
WHERE transaction_id = $1 LIMIT 1;

-- name: UpdateAccountTransaction :one
UPDATE account_transactions 
SET transaction_status = $2
WHERE transaction_id = $1
RETURNING *;

-- name: ListAccountTransactions :many
SELECT * FROM account_transactions
WHERE account_id = $1 
ORDER BY transaction_id
LIMIT $2 OFFSET $3;

