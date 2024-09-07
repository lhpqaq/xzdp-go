package redis

import (
	"context"
	"log"
	"time"
)

type Lock struct {
	Ctx    context.Context
	Key    string
	Value  string
	Expire time.Duration
}

func NewLock(ctx context.Context, key string, value string, expire time.Duration) *Lock {
	return &Lock{
		Ctx:    ctx,
		Key:    key,
		Value:  value,
		Expire: expire,
	}
}

func (l *Lock) TryLock() bool {
	// value 应当全局唯一
	success, err := RedisClient.SetNX(l.Ctx, l.Key, l.Value, l.Expire*time.Second).Result()
	if err != nil {
		log.Printf("Error acquiring lock: %v", err)
		return false
	}
	return success
}

func (l *Lock) UnLock(value string) {
	luaScript := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        end
		return 0
    `
	RedisClient.Eval(l.Ctx, luaScript, []string{l.Key}, value)
}
