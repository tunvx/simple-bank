package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
)

// GetCacheAccount method fetches account data from cache
func (r *RedisCache) GetCacheAccount(ctx context.Context, accountNumber string) (db.Account, error) {
	cacheKey := fmt.Sprintf("account:%s", accountNumber)
	result, err := r.client.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return db.Account{}, nil
	}
	if err != nil {
		return db.Account{}, err
	}

	// Parse JSON thành db.Account
	var account db.Account
	err = json.Unmarshal([]byte(result), &account)
	if err != nil {
		return db.Account{}, err
	}
	return account, nil
}

// SetCacheAccount method stores account data to cache
func (r *RedisCache) SetCacheAccount(ctx context.Context, account db.Account) error {
	cacheKey := fmt.Sprintf("account:%s", account.AccountNumber)

	// Chuyển đổi account struct sang JSON
	accountJSON, err := json.Marshal(account)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal account")
		return err
	}

	// Lưu JSON vào Redis
	err = r.client.Set(ctx, cacheKey, accountJSON, accountCacheExpiration).Err()
	return err
}

// DeleteCacheAccount method deletes account data from cache
func (r *RedisCache) DeleteCacheAccount(ctx context.Context, accountNumber string) error {
	cacheKey := fmt.Sprintf("account:%s", accountNumber)
	err := r.client.Del(ctx, cacheKey).Err()
	return err
}
