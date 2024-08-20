package user

import (
	"context"
	"testing"

	user "xzdp/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestUserLoginService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUserLoginService(ctx, c)
	// init req and assert value
	req := &user.UserLoginFrom{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
