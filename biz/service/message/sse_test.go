package message

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	commonTestAssert "github.com/cloudwego/hertz/pkg/common/test/assert"
	goredis "github.com/go-redis/redis/v8"
	
	"xzdp/biz/dal/redis"
	"xzdp/biz/model/user"
	"xzdp/biz/utils"
)

func TestSseService_Run(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer mr.Close()
	redis.RedisClient = goredis.NewClient(&goredis.Options{Addr: mr.Addr()})

	ctx, cancel := context.WithCancel(context.Background())
	c := app.NewContext(1)
	
	// Inject User
	u := &user.UserDTO{ID: 1}
	ctx = utils.SaveUser(ctx, u)

	s := NewSseService(ctx, c)
	req := ">"
	
	// Cancel the context immediately to prevent infinite loop
	cancel()
	
	resp, err := s.Run(req)
	// It's acceptable to have an error if context is canceled, 
	// or nil if handled gracefully. The goal is no panic and prompt return.
	commonTestAssert.DeepEqual(t, (*string)(nil), resp)
	if err != nil && err != context.Canceled {
		t.Errorf("Unexpected error: %v", err)
	}
}