-- name: InsertAccountShardMap :one
INSERT INTO account_shard_map (
  account_id,
  customer_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetShardByAccountID :one
SELECT c.shard_id FROM account_shard_map a
JOIN customer_shard_map c ON a.customer_id = c.customer_id
WHERE a.account_id = $1;


