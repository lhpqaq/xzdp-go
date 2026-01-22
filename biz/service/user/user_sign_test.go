package user

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/go-redis/redis/v8"

	model "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"
)

func TestUserSignService_Run(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mr.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	ctx := context.Background()
	c := app.NewContext(1)

	// Case 1: Success
	t.Run("Success", func(t *testing.T) {
		userId := int64(1001)
		expectedUser := &model.UserDTO{
			ID: userId,
		}
		ctxWithUser := utils.SaveUser(ctx, expectedUser)

		s := NewUserSignServiceWithRedis(ctxWithUser, c, rdb)

		req := &model.Empty{}
		resp, err := s.Run(req)

		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, true, *resp)

		// Verify Redis
		now := time.Now()
		keySuffix := now.Format(":200601")
		key := constants.USER_SIGN_KEY + fmt.Sprint(userId) + keySuffix
		dayOfMonth := now.Day()

		bit, _ := rdb.GetBit(ctx, key, int64(dayOfMonth-1)).Result()
		assert.DeepEqual(t, int64(1), bit)
	})

	// Case 2: Not Logged In
	t.Run("Fail_NotLoggedIn", func(t *testing.T) {
		s := NewUserSignServiceWithRedis(ctx, c, rdb) // empty context

		req := &model.Empty{}
		resp, err := s.Run(req)

		assert.NotEqual(t, nil, err)
		assert.DeepEqual(t, "用户未登录", err.Error())
		assert.DeepEqual(t, (*bool)(nil), resp)
	})
}
