package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
	"xzdp/biz/model/cache"
	"xzdp/biz/pkg/constants"
	"xzdp/conf"
)

var RedisClient *redis.Client

type ArgsFunc func(args ...interface{}) (interface{}, error)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Password: conf.GetConf().Redis.Password,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

func SetString(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = RedisClient.Set(ctx, key, jsonData, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetStringLogical(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	stringValue, err := json.Marshal(value)
	cacheData := cache.NewRedisStringData(string(stringValue), time.Now().Add(duration))
	jsonData, err := json.Marshal(*cacheData)
	if err != nil {
		return err
	}

	err = RedisClient.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

//func GetString(ctx context.Context, key string, duration time.Duration) (interface{}, error) {
//
//}

func GetStringLogical(ctx context.Context, key string, duration time.Duration, dbFallback ArgsFunc, args ...interface{}) (string, error) {
	redisJson, err := RedisClient.Get(ctx, key).Result()
	lock := NewLock(ctx, constants.LOCK_KEY+key, "lock", duration)
	if redisJson == "" || errors.Is(err, redis.Nil) {
		if lock.TryLock() {
			go func() {
				defer lock.UnLock("lock")
				data, err := dbFallback(args...)
				if err != nil {
					return
				}
				err = SetStringLogical(ctx, key, data, duration)
				if err != nil {
					return
				}
			}()
		}
	} else {
		var redisData cache.RedisStringData
		if err := json.Unmarshal([]byte(redisJson), &redisData); err != nil {
			return "", err
		}
		if redisData.ExpiredTime.After(time.Now()) {
			return redisData.Data, nil
		} else {
			if lock.TryLock() {
				go func() {
					defer lock.UnLock("lock")
					data, err := dbFallback(args...)
					if err != nil {
						return
					}
					err = SetStringLogical(ctx, key, data, duration)
					if err != nil {
						return
					}
				}()
			}
			return redisData.Data, nil
		}
	}
	return "", nil
}
