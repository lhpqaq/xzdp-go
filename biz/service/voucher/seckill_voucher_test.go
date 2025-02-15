package voucher

import (
	"context"
	"fmt"
	"testing"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	"xzdp/biz/model/user"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func setup() {
	mysql.Init()
	redis.Init()
}

func TestSeckillVoucherService_Run(t *testing.T) {
	setup()
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewSeckillVoucherService(ctx, c)
	s.Context = utils.SaveUser(s.Context, &user.UserDTO{
		ID:       6,
		NickName: "test",
	})
	// init req and assert value
	req := int64(1)
	resp, err := s.Run(&req)
	fmt.Println(resp, err)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}

func TestSeckillVoucherService_StressTest(t *testing.T) {
	setup()

	concurrency := 10
	requests := 100

	done := make(chan bool, concurrency)

	for i := 0; i < requests; i++ {
		go func(i int) {
			defer func() { done <- true }()
			ctx := context.Background()
			c := app.NewContext(1)
			s := NewSeckillVoucherService(ctx, c)
			s.Context = utils.SaveUser(s.Context, &user.UserDTO{
				ID:       int64(i),
				NickName: "test_" + fmt.Sprintf("%d", i),
			})
			req := int64(1)
			resp, err := s.Run(&req)
			if err != nil {
				t.Errorf("request %d failed: %v", i, err)
			}
			if resp != nil {
				t.Errorf("request %d got unexpected response: %v", i, resp)
			}
		}(i)
	}

	// 等待所有请求完成
	for i := 0; i < requests; i++ {
		<-done
	}
}
