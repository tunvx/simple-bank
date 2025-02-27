package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	accountCacheExpiration = time.Minute * 10
)

type Cache interface {
	GetCacheAccountInfo(ctx context.Context, key int64) (*AccountInfo, error)
	SetCacheAccountInfo(ctx context.Context, value *AccountInfo) error
	DeleteCacheAccountInfo(ctx context.Context, key int64) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(redisOpts *redis.Options) Cache {
	client := redis.NewClient(redisOpts)
	return &RedisCache{client: client}
}
