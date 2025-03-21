// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: customer_shard_map.sql

package db

import (
	"context"
)

const getShardByCustomeRid = `-- name: GetShardByCustomeRid :one
SELECT customer_id, customer_rid, shard_id FROM customer_shard_map
WHERE customer_rid = $1
`

func (q *Queries) GetShardByCustomeRid(ctx context.Context, customerRid string) (CustomerShardMap, error) {
	row := q.db.QueryRow(ctx, getShardByCustomeRid, customerRid)
	var i CustomerShardMap
	err := row.Scan(&i.CustomerID, &i.CustomerRid, &i.ShardID)
	return i, err
}

const insertCustomerShardMap = `-- name: InsertCustomerShardMap :one
INSERT INTO customer_shard_map (
  customer_rid,
  shard_id
) VALUES (
  $1, $2
) RETURNING customer_id, customer_rid, shard_id
`

type InsertCustomerShardMapParams struct {
	CustomerRid string `json:"customer_rid"`
	ShardID     int32  `json:"shard_id"`
}

func (q *Queries) InsertCustomerShardMap(ctx context.Context, arg InsertCustomerShardMapParams) (CustomerShardMap, error) {
	row := q.db.QueryRow(ctx, insertCustomerShardMap, arg.CustomerRid, arg.ShardID)
	var i CustomerShardMap
	err := row.Scan(&i.CustomerID, &i.CustomerRid, &i.ShardID)
	return i, err
}

const updateCustomerShardMap = `-- name: UpdateCustomerShardMap :one
UPDATE customer_shard_map
SET shard_id = $2
WHERE customer_id = $1
RETURNING customer_id, customer_rid, shard_id
`

type UpdateCustomerShardMapParams struct {
	CustomerID int64 `json:"customer_id"`
	ShardID    int32 `json:"shard_id"`
}

func (q *Queries) UpdateCustomerShardMap(ctx context.Context, arg UpdateCustomerShardMapParams) (CustomerShardMap, error) {
	row := q.db.QueryRow(ctx, updateCustomerShardMap, arg.CustomerID, arg.ShardID)
	var i CustomerShardMap
	err := row.Scan(&i.CustomerID, &i.CustomerRid, &i.ShardID)
	return i, err
}
