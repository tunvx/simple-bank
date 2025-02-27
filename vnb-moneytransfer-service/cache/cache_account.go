package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
)

type AccountInfo struct {
	AccountID     int64
	ShardId       int
	CustomerID    int64
	OwnerName     string
	CurrencyType  db.Currencytype
	AccountStatus db.Accountstatus
}

// GetCacheAccountInfo method fetches account data from cache
func (r *RedisCache) GetCacheAccountInfo(ctx context.Context, accountId int64) (*AccountInfo, error) {
	cacheKey := fmt.Sprintf("account:%d", accountId)
	log.Debug().Msgf("Cache: Fetching account ( %d )", accountId)

	result, err := r.client.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		log.Debug().Msgf("Cache: MISS - account ( %d ) not found in cache", accountId)
		return nil, nil
	}
	if err != nil {
		log.Error().Err(err).Msgf("Cache: Failed to get account ( %d ) from cache", accountId)
		return nil, err
	}

	// Parse JSON thành db.Account
	var accountInfo AccountInfo
	err = json.Unmarshal([]byte(result), &accountInfo)
	if err != nil {
		log.Error().Err(err).Msgf("Cache: Failed to unmarshal account ( %d ) from cache", accountId)
		return nil, err
	}

	log.Debug().Msgf("Cache: HIT - Successfully fetched account ( %d )", accountId)
	return &accountInfo, nil
}

// SetCacheAccountInfo method stores account data to cache
func (r *RedisCache) SetCacheAccountInfo(ctx context.Context, account *AccountInfo) error {
	cacheKey := fmt.Sprintf("account:%d", account.AccountID)
	log.Debug().Msgf("Cache: Storing account ( %d )", account.AccountID)

	// Chuyển đổi account struct sang JSON
	accountJSON, err := json.Marshal(account)
	if err != nil {
		log.Error().Err(err).Msg("Cache: Failed to marshal account")
		return err
	}

	// Lưu JSON vào Redis
	err = r.client.Set(ctx, cacheKey, accountJSON, accountCacheExpiration).Err()
	if err != nil {
		log.Error().Err(err).Msgf("Cache: Failed to store account ( %d )", account.AccountID)
		return err
	}

	log.Debug().Msgf("Cache: Successfully stored account ( %d )", account.AccountID)
	return nil
}

// DeleteCacheAccountInfo method deletes account data from cache
func (r *RedisCache) DeleteCacheAccountInfo(ctx context.Context, accountId int64) error {
	cacheKey := fmt.Sprintf("account:%d", accountId)
	log.Debug().Msgf("Cache: Deleting account ( %d )", accountId)

	err := r.client.Del(ctx, cacheKey).Err()
	if err != nil {
		log.Error().Err(err).Msgf("Cache: Failed to delete account ( %d )", accountId)
		return err
	}

	log.Debug().Msgf("Cache: Successfully deleted account ( %d )", accountId)
	return nil
}
