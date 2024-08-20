package shop

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	shop "xzdp/biz/model/shop"
)

func TestShopOfTypeService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewShopOfTypeService(ctx, c)
	// init req and assert value
	req := &shop.ShopOfTypeReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
