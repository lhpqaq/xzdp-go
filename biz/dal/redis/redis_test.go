package redis

import (
	"context"
	"fmt"
	"log"
	"math"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestHyperLogLog(t *testing.T) {
	Init()
	values := make([]interface{}, 1000)
	ctx := context.Background()
	var total int64 = 1000000
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
	diff := math.Abs(float64(total - count))
	assert.True(t, count > 0)
	assert.True(t, diff < float64(total)/100)
}
