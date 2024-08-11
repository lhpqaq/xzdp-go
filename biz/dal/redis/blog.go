package redis

import (
	"context"
	"errors"
	redis2 "github.com/go-redis/redis/v8"
	"xzdp/biz/pkg/constants"
)

func IsLiked(ctx context.Context, key string, member string) (ok bool, err error) {
	_, err = RedisClient.ZScore(ctx, key, member).Result()
	if err != nil {
		if errors.Is(err, redis2.Nil) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func HasLikes(ctx context.Context, key string) (bool, error) {
	exists, err := RedisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func DeleteLikes(ctx context.Context, key string) error {
	err := RedisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetBlogsByKey(ctx context.Context, key string, max string, offset int64) ([]redis2.Z, error) {
	size := constants.MAX_PAGE_SIZE
	result, err := RedisClient.ZRevRangeByScoreWithScores(
		ctx,
		key,
		&redis2.ZRangeBy{Max: max, Min: "0", Offset: offset, Count: int64(size)},
	).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
