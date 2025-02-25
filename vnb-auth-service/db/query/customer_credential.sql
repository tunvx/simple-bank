-- name: CreateCustomerCredential :one
INSERT INTO customer_credentials (
  customer_id,
  shard_id,
  username,
  hashed_password
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetCustomerCredential :one
SELECT * FROM customer_credentials
WHERE username = $1 LIMIT 1;

-- name: UpdateCustomerCredential :one
UPDATE customer_credentials
SET 
  username = COALESCE(sqlc.narg(username), username),
  username_changed_at = CASE
    WHEN sqlc.narg(username) IS NOT NULL THEN now()
    ELSE username_changed_at
  END,
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = CASE
    WHEN sqlc.narg(hashed_password) IS NOT NULL THEN now()
    ELSE password_changed_at
  END,
  shard_id = COALESCE(sqlc.narg(shard_id), shard_id),
  shard_id_changed_at = CASE
    WHEN sqlc.narg(shard_id) IS NOT NULL THEN now()
    ELSE shard_id_changed_at
  END
WHERE
  customer_id = sqlc.arg(customer_id)
RETURNING *;

-- -- name: UpdateCustomerCredential :one
-- UPDATE customer_credentials
-- SET 
--   username = CASE
--     WHEN @set_username::boolean = TRUE THEN @username
--     ELSE username
--   END,
--   username_changed_at = CASE
--     WHEN @set_username::boolean = TRUE THEN now()
--     ELSE username_changed_at
--   END,
--   hashed_password = CASE
--     WHEN @set_hashed_password::boolean = TRUE THEN @hashed_password
--     ELSE hashed_password
--   END,
--   password_changed_at = CASE
--     WHEN @set_hashed_password::boolean = TRUE THEN now()
--     ELSE password_changed_at
--   END
-- WHERE
--   customer_id = @customer_id
-- RETURNING *;