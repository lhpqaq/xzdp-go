package redis

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"
	"xzdp/biz/model/shop"

	"github.com/go-redis/redis/v8"
)

func GetShopFromCache(ctx context.Context, key string) (*shop.Shop, error) {
	shopJson, err := RedisClient.Get(ctx, key).Result()
	if err == nil && shopJson != "" {
		var shop shop.Shop
		if err := json.Unmarshal([]byte(shopJson), &shop); err != nil {
			return nil, err
		}
		return &shop, nil
	}

	if err != redis.Nil {
		return nil, err
	}

	if shopJson == "" {
		return nil, errors.New("shop not found in cache")
	}

	return nil, errors.New("unknown error")
}

func TryLock(ctx context.Context, key string) bool {
	success, err := RedisClient.SetNX(ctx, key, "1", 10*time.Second).Result()
	if err != nil {
		log.Printf("Error acquiring lock: %v", err)
		return false
	}
	return success
}

func UnLock(ctx context.Context, key string) {
	RedisClient.Del(ctx, key)
}
