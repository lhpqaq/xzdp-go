package redis

import (
	"context"
	"encoding/json"
	"errors"
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
