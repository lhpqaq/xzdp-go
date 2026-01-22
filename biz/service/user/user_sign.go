package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"xzdp/biz/dal/redis"
	user "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	red "github.com/go-redis/redis/v8"
)

type UserSignService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	Redis          *red.Client
}

func NewUserSignService(ctx context.Context, requestContext *app.RequestContext) *UserSignService {
	return &UserSignService{
		RequestContext: requestContext,
		Context:        ctx,
		Redis:          redis.RedisClient,
	}
}

// NewUserSignServiceWithRedis for testing
func NewUserSignServiceWithRedis(ctx context.Context, requestContext *app.RequestContext, rds *red.Client) *UserSignService {
	return &UserSignService{
		RequestContext: requestContext,
		Context:        ctx,
		Redis:          rds,
	}
}

func (h *UserSignService) Run(req *user.Empty) (resp *bool, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	userdto := utils.GetUser(h.Context)
	if userdto == nil {
		return nil, errors.New("用户未登录")
	}
	userId := userdto.ID
	now := time.Now()
	keySuffix := now.Format(":200601")
	key := constants.USER_SIGN_KEY + fmt.Sprint(userId) + keySuffix
	dayOfMonth := now.Day()
	err = h.Redis.SetBit(h.Context, key, int64(dayOfMonth-1), 1).Err()
	if err != nil {
		return nil, err
	}
	boolResp := true
	return &boolResp, nil
}
