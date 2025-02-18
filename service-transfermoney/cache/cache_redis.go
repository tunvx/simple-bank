package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	db "github.com/tunvx/simplebank/management/db/sqlc"
)

const (
	accountCacheExpiration = time.Minute * 10
)

type Cache interface {
	GetCacheAccount(ctx context.Context, key string) (db.Account, error)
	SetCacheAccount(ctx context.Context, value db.Account) error
	DeleteCacheAccount(ctx context.Context, key string) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(redisOpts *redis.Options) *RedisCache {
	client := redis.NewClient(redisOpts)
	return &RedisCache{client: client}
}
