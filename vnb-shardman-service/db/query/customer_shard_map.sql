-- name: InsertCustomerShardMap :one
INSERT INTO customer_shard_map (
  customer_rid,
  shard_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetShardByCustomeRid :one
SELECT * FROM customer_shard_map
WHERE customer_rid = $1;

-- name: UpdateCustomerShardMap :one
UPDATE customer_shard_map
SET shard_id = $2
WHERE customer_id = $1
RETURNING *;