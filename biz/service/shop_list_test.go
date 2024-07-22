package service

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	xzdp "xzdp/biz/model/xzdp"
)

func TestShopListService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewShopListService(ctx, c)
	// init req and assert value
	req := &xzdp.Empty{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
