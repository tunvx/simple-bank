// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: customer_credential.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCustomerCredential = `-- name: CreateCustomerCredential :one
INSERT INTO customer_credentials (
  customer_id,
  username,
  hashed_password
) VALUES (
  $1, $2, $3
) RETURNING customer_id, username, hashed_password, created_at, username_changed_at, password_changed_at
`

type CreateCustomerCredentialParams struct {
	CustomerID     int64  `json:"customer_id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateCustomerCredential(ctx context.Context, arg CreateCustomerCredentialParams) (CustomerCredential, error) {
	row := q.db.QueryRow(ctx, createCustomerCredential, arg.CustomerID, arg.Username, arg.HashedPassword)
	var i CustomerCredential
	err := row.Scan(
		&i.CustomerID,
		&i.Username,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UsernameChangedAt,
		&i.PasswordChangedAt,
	)
	return i, err
}

const getCustomerCredential = `-- name: GetCustomerCredential :one
SELECT customer_id, username, hashed_password, created_at, username_changed_at, password_changed_at FROM customer_credentials
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetCustomerCredential(ctx context.Context, username string) (CustomerCredential, error) {
	row := q.db.QueryRow(ctx, getCustomerCredential, username)
	var i CustomerCredential
	err := row.Scan(
		&i.CustomerID,
		&i.Username,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UsernameChangedAt,
		&i.PasswordChangedAt,
	)
	return i, err
}

const updateCustomerCredential = `-- name: UpdateCustomerCredential :one
UPDATE customer_credentials
SET 
  username = COALESCE($1, username),
  username_changed_at = CASE
    WHEN $1 IS NOT NULL THEN now()
    ELSE username_changed_at
  END,
  hashed_password = COALESCE($2, hashed_password),
  password_changed_at = CASE
    WHEN $2 IS NOT NULL THEN now()
    ELSE password_changed_at
  END
WHERE
  customer_id = $3
RETURNING customer_id, username, hashed_password, created_at, username_changed_at, password_changed_at
`

type UpdateCustomerCredentialParams struct {
	Username       pgtype.Text `json:"username"`
	HashedPassword pgtype.Text `json:"hashed_password"`
	CustomerID     int64       `json:"customer_id"`
}

func (q *Queries) UpdateCustomerCredential(ctx context.Context, arg UpdateCustomerCredentialParams) (CustomerCredential, error) {
	row := q.db.QueryRow(ctx, updateCustomerCredential, arg.Username, arg.HashedPassword, arg.CustomerID)
	var i CustomerCredential
	err := row.Scan(
		&i.CustomerID,
		&i.Username,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UsernameChangedAt,
		&i.PasswordChangedAt,
	)
	return i, err
}