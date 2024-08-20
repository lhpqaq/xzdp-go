package shop

import (
	"context"
	"sync"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestShopInfoService_Run(t *testing.T) {
	var shopId int64 = 1
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewShopInfoService(ctx, c)
	resp, err := s.Run(shopId)
	assert.NotEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	refShop := resp
	var wg sync.WaitGroup
	getShopInfo := func() {
		defer wg.Done()
		ctx := context.Background()
		c := app.NewContext(1)
		s := NewShopInfoService(ctx, c)
		for i := 0; i < 100; i++ {
			resp, err := s.Run(shopId)
			assert.DeepEqual(t, refShop, resp)
			assert.DeepEqual(t, nil, err)
		}
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getShopInfo()
	}
	wg.Wait()
	// todo edit your unit test.
}
