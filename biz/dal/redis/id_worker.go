package redis

import (
	"context"
	"time"
	"xzdp/biz/pkg/constants"
)

func NextId(ctx context.Context, prefix string) (int64, error) {
	now := time.Now().Unix()
	dateKey := time.Now().Format("2006-01-02")

	count, err := RedisClient.Incr(ctx, constants.ID_KEY+prefix+":"+dateKey).Result()
	if err != nil {
		return 0, err
	}
	return (now << 32) | count, nil
}
