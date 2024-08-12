package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"
)

// ProduceMq 写stream消息队列
func ProduceMq(ctx context.Context, key string, message interface{}) error {
	messageJSON, err := utils.SerializeStruct(message)
	if err != nil {
		return err
	}
	// xadd写stream
	err = RedisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: key,
		ID:     "*",
		Values: []interface{}{"message", messageJSON},
	}).Err()
	// 错误处理
	if err != nil {
		return err
	}
	return nil
}

// ConsumeMq 读取stream消息队列
func ConsumeMq(ctx context.Context, key string, consumer string, block time.Duration, count int64, id string) ([]redis.XStream, error) {
	if id == "" {
		id = ">"
	}
	xSet, err := RedisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    constants.STREAM_READ_GROUP,
		Consumer: consumer,
		Streams:  []string{key, id},
		Count:    count,
		Block:    block,
		NoAck:    false,
	}).Result()
	if err != nil {
		return nil, err
	}
	return xSet, nil
}
func CreateConsumerGroup(ctx context.Context, key string) {
	RedisClient.XGroupCreateMkStream(ctx, key, constants.STREAM_READ_GROUP, "0")
}
func AckMq(ctx context.Context, key string, id string) error {
	return RedisClient.XAck(ctx, key, constants.STREAM_READ_GROUP, id).Err()
}
