package redis

import (
	"context"
	"sync"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

func TestNextId(t *testing.T) {
	// Mock Redis
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer mr.Close()
	RedisClient = redis.NewClient(&redis.Options{Addr: mr.Addr()})

	// 测试用例
	testCases := []struct {
		name    string
		prefix  string
		wantErr bool
	}{
		{"valid_prefix", "test", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var idSet sync.Map // 使用sync.Map来存储ID
			var wg sync.WaitGroup
			numGoroutines := 100 // Reduce for faster test in CI

			// 启动协程
			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for i := 0; i < 50; i++ {
						id, err := NextId(context.Background(), tc.prefix)
						if err != nil {
							t.Errorf("NextId() returned error: %v", err)
							return
						}
						if id <= 0 {
							t.Errorf("NextId() returned invalid ID: %d", id)
						}
						idSet.Store(id, struct{}{}) // 将ID放入sync.Map中
					}
				}()
			}

			// 等待所有协程完成
			wg.Wait()

			// 检查生成的ID数量
			var uniqueCount int
			idSet.Range(func(key, value interface{}) bool {
				uniqueCount++
				return true // 返回true继续遍历
			})
			t.Log("uniqueCount:", uniqueCount)
			if uniqueCount != numGoroutines*50 {
				t.Errorf("Expected %d unique IDs, got %d", numGoroutines*50, uniqueCount)
			}
		})
	}
}
