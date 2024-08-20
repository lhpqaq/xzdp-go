package shop

import (
	"context"
	"testing"

	shop "xzdp/biz/model/shop"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestShopListService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewShopListService(ctx, c)
	// init req and assert value
	req := &shop.Empty{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
