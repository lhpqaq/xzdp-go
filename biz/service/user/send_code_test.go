package user

import (
	"context"
	"testing"

	redis2 "xzdp/biz/dal/redis"
	user "xzdp/biz/model/user"

	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/go-redis/redis/v8"
)

func TestSendCodeService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewSendCodeService(ctx, c)
	// init req and assert value
	req := &user.UserLoginFrom{
		Phone: "12345678901"}
	resp, err := s.Run(req)
	assert.Nil(t, resp)
	assert.NotNil(t, err)

	resp, err = s.Run(&user.UserLoginFrom{
		Phone: "13412332123"})

	assert.DeepEqual(t, &user.Result{Success: true}, resp)
	assert.DeepEqual(t, nil, err)
}

func TestMain(m *testing.M) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	redis2.RedisClient = redis.NewClient(&redis.Options{
		Addr: s.Addr(), // mock redis server的地址
	})
	m.Run()
}
