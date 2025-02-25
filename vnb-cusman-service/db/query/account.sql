-- name: CreateAccount :one
INSERT INTO accounts (
    account_id,
    customer_id,
    current_balance,
    currency_type,
    account_status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE account_id = $1
LIMIT 1;

-- name: CheckAccountByID :one
SELECT 
  a.account_id, 
  c.full_name AS customer_name, 
  a.currency_type, 
  a.account_status
FROM accounts a
JOIN customers c ON a.customer_id = c.customer_id
WHERE a.account_id = sqlc.arg(account_id)
LIMIT 1;

-- name: ListAccountByCustomerID :many
SELECT * FROM accounts
WHERE customer_id = $1
ORDER BY account_id;


-- The methods below are directly related to transactions (multiple queries), so will be implemented query with id
-- Especially directly related to the locking mechanism, locking in order in the database
-- ...

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateAccountBalance :one
UPDATE accounts 
SET current_balance = $2
WHERE account_id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
SET current_balance = current_balance + sqlc.arg(amount)
WHERE account_id = sqlc.arg(account_id)
RETURNING *;