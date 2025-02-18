-- name: CreateCustomerSession :one
INSERT INTO customer_sessions (
  session_id,
  customer_id,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetCustomerSession :one
SELECT * FROM customer_sessions
WHERE session_id = $1 LIMIT 1;

-- noti-tag