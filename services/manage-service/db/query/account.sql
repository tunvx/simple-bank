-- name: CreateAccount :one
INSERT INTO accounts (
    account_number,
    customer_id,
    current_balance,
    currency_type,
    account_status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListAccountByCustomerID :many
SELECT * FROM accounts
WHERE 
    customer_id = $1 AND
    account_status = $2
ORDER BY account_id
LIMIT $3
OFFSET $4;

-- name: GetAccountByAccNumber :one
SELECT * FROM accounts
WHERE account_number = $1
LIMIT 1;

-- name: GetAccountIDByAccNumber :one
SELECT account_id FROM accounts
WHERE account_number = $1
LIMIT 1;

-- name: GetCustomerIDByAccNumber :one
SELECT customer_id FROM accounts
WHERE account_number = $1
LIMIT 1;

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