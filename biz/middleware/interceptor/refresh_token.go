package interceptor

import (
	"context"
	"time"
	"xzdp/biz/dal/redis"
	model "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func CheckToken(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "check token interceptor")
	token := c.GetHeader("authorization")
	if token == nil {
		c.Next(ctx)
	}
	if len(token) == 0 {
		c.Next(ctx)
	}
	var userdto model.UserDTO
	if err := redis.RedisClient.HGetAll(ctx, constants.LOGIN_USER_KEY+string(token)).Scan(&userdto); err != nil {
		c.Next(ctx)
	}
	redis.RedisClient.Expire(ctx, constants.LOGIN_USER_KEY+string(token), time.Minute*1)
	ctx = utils.SaveUser(ctx, &userdto)
	c.Next(ctx)
	if utils.GetUser(ctx) == nil {
		hlog.CtxErrorf(ctx, "check token interceptor error")
	}
	hlog.CtxDebugf(ctx, "user = %+v", utils.GetUser(ctx))
}

func LoginInterceptor(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "login interceptor")
	if utils.GetUser(ctx) == nil {
		c.SetStatusCode(401)
		c.Abort()
	}
	c.Next(ctx)
}
