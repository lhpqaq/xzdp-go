package service

import (
	"context"
	"testing"
	model "xzdp/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestSendCodeService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewSendCodeService(ctx, c)
	// init req and assert value
	req := &model.UserLoginFrom{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
