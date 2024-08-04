package service

import (
	"context"
	"testing"

	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"

	user "xzdp/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestSendCodeService_Run(t *testing.T) {
	// redis.Init()

	ctx := context.Background()
	c := app.NewContext(1)
	s := NewSendCodeService(ctx, c)
	// init req and assert value
	req := &user.UserLoginFrom{
		Phone: "12345678901"}
	resp, err := s.Run(req)

	assert.DeepEqual(t, &user.Result{Success: true}, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}

func TestMain(m *testing.M) {
	redis.Init()
	mysql.Init()
	m.Run()
}
