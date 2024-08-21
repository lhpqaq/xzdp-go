package redis

import (
	"context"
	"log"
	"time"
)

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
