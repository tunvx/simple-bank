-- name: InsertCustomerShardMap :one
INSERT INTO customer_shard_map (
  customer_rid,
  shard_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetShardByCustomeRid :one
SELECT * FROM customer_shard_map
WHERE customer_rid = $1 LIMIT 1;
