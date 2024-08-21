package cache

import "time"

type RedisStringData struct {
	Data        string
	ExpiredTime time.Time
}

func NewRedisStringData(data string, expiredTime time.Time) *RedisStringData {
	return &RedisStringData{data, expiredTime}
}
