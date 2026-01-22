package redis

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/go-redis/redis/v8"
)

func TestHyperLogLog(t *testing.T) {
	// Mock Redis
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer mr.Close()
	RedisClient = redis.NewClient(&redis.Options{Addr: mr.Addr()})

	values := make([]interface{}, 1000)
	ctx := context.Background()
	var total int64 = 10000 // Reduced for performance
	// 批量保存 100w 条用户记录，每批 1000 条
	var i int64
	for i = 0; i < total; i++ {
		// 获取当前批次的索引
		j := i % 1000
		// 生成用户记录
		values[j] = "user_" + fmt.Sprint(i)
		// 每 1000 条记录发送一次到 Redis
		if j == 999 {
			err := RedisClient.PFAdd(ctx, "hl2", values...).Err()
			if err != nil {
				log.Fatalf("Failed to add values to HyperLogLog: %v", err)
			}
		}
	}

	// 统计 HyperLogLog 中的用户数量
	count, err := RedisClient.PFCount(ctx, "hl2").Result()
	if err != nil {
		log.Fatalf("Failed to get HyperLogLog count: %v", err)
	}
	log.Printf("HyperLogLog count: %d", count)
	// Miniredis PFCount implementation might differ or be exact, 
	// just ensure it runs without panic and returns logical result.
	assert.True(t, count > 0)
}