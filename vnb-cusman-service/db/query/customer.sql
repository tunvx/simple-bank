-- name: CreateCustomer :one
INSERT INTO customers (
  customer_id,
  customer_rid,
  full_name,
  date_of_birth,
  permanent_address,
  phone_number,
  email_address,
  customer_tier,
  customer_segment,
  financial_status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: GetCustomerByID :one
SELECT * FROM customers
WHERE customer_id = $1 LIMIT 1;

-- name: GetCustomerByRid :one
SELECT * FROM customers
WHERE customer_rid = $1 LIMIT 1;

-- name: UpdateCustomer :one
UPDATE customers
SET
  phone_number = COALESCE(sqlc.narg(phone_number), phone_number),
  email_address = COALESCE(sqlc.narg(email_address), email_address),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
  customer_id = sqlc.arg(customer_id) -- id from auth_token
RETURNING *;