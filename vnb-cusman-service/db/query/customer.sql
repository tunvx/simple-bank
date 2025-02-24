-- name: CreateCustomer :one
INSERT INTO customers (
  customer_rid,
  fullname,
  date_of_birth,
  address,
  phone_number,
  email,
  customer_tier,
  customer_segment,
  financial_status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetCustomerByID :one
SELECT * FROM customers
WHERE customer_id = $1 LIMIT 1;

-- name: GetCustomerByRid :one
SELECT * FROM customers
WHERE customer_rid = $1 LIMIT 1;

-- name: GetCustomerIdByRid :one
SELECT customer_id FROM customers
WHERE customer_rid = $1 LIMIT 1;

-- name: GetCustomerNameByID :one
SELECT fullname FROM customers
WHERE customer_id = $1 LIMIT 1;

-- name: UpdateCustomer :one
UPDATE customers
SET
  address = COALESCE(sqlc.narg(address), address),
  phone_number = COALESCE(sqlc.narg(phone_number), phone_number),
  email = COALESCE(sqlc.narg(email), email),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
  customer_id = sqlc.arg(customer_id)
RETURNING *;