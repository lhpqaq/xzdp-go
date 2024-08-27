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
)

type UserSignService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserSignService(Context context.Context, RequestContext *app.RequestContext) *UserSignService {
	return &UserSignService{RequestContext: RequestContext, Context: Context}
}

func (h *UserSignService) Run(req *user.Empty) (resp *bool, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	userdto := utils.GetUser(h.Context)
	if userdto == nil {
		return nil, errors.New("用户未登录")
	}
	userId := userdto.ID
	now := time.Now()
	keySuffix := now.Format(":200601")
	key := constants.USER_SIGN_KEY + fmt.Sprint(userId) + keySuffix
	dayOfMonth := now.Day()
	err = redis.RedisClient.SetBit(h.Context, key, int64(dayOfMonth-1), 1).Err()
	if err != nil {
		return nil, err
	}
	boolResp := true
	return &boolResp, nil
}
