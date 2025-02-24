-- name: InsertAccountShardMap :one
INSERT INTO account_shard_map (
  account_id,
  customer_id,
  shard_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetShardByAccountID :one
SELECT * FROM account_shard_map
WHERE account_id = $1 LIMIT 1;
