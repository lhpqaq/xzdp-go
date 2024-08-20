package xzdp

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	xzdp "xzdp/biz/model/xzdp"
)

func TestHelloMethodService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewHelloMethodService(ctx, c)
	// init req and assert value
	req := &xzdp.HelloReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
