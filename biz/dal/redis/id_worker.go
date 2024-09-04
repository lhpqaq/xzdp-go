package redis

import (
	"context"
	"time"
	"xzdp/biz/pkg/constants"
)

func NextId(ctx context.Context, prefix string) (int64, error) {
	beginTimeStamp := int64(constants.BEGIN_TIMESTAMP)
	now := time.Now().Unix()
	timeStamp := now - beginTimeStamp
	dateKey := time.Unix(now, 0).Format("2006-01-02")
	count, err := RedisClient.Incr(ctx, constants.ICRID_KEY+prefix+":"+dateKey).Result()
	if err != nil {
		return 0, err
	}
	return (timeStamp << constants.COUNT_BIT) | count, nil
}
