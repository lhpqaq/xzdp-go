package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	redis2 "github.com/go-redis/redis/v8"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	"xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"
)

func GetUserDtoFromCacheOrDB(ctx context.Context, userID int64) (*user.UserDTO, error) {
	key := fmt.Sprintf("%s%d", constants.CACHE_USERDTO_KEY, userID)
	// Try to get the user from cache.
	cachedUser, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis2.Nil) {
		return nil, err
	}
	if err == nil {
		// If we have a value, unmarshal it into a user struct.
		var interUser user.User
		if err := json.Unmarshal([]byte(cachedUser), &interUser); err != nil {
			return nil, err
		}
		dto := utils.UserToUserDTO(&interUser)
		return dto, nil
	}

	// If not in cache, fetch from DB.
	interUser, err := mysql.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Save the user to cache.
	userJSON, err := json.Marshal(interUser)
	if err != nil {
		return nil, err
	}
	if err := redis.RedisClient.Set(ctx, key, userJSON, constants.CACHE_USERDTO_EXPIRE).Err(); err != nil {
		hlog.CtxWarnf(ctx, "Failed to set user cache: %v", err)
	}
	dto := utils.UserToUserDTO(interUser)
	return dto, nil
}
