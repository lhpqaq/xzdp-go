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
)

type UserSignCountService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserSignCountService(Context context.Context, RequestContext *app.RequestContext) *UserSignCountService {
	return &UserSignCountService{RequestContext: RequestContext, Context: Context}
}

func (h *UserSignCountService) Run(req *user.Empty) (resp *int64, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
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
	var signCount int64 = 0
	bitFieldResult, err := redis.RedisClient.BitField(h.Context, key, "GET", fmt.Sprintf("u%d", dayOfMonth), "0").Result()
	if err != nil {
		return nil, err
	}
	if len(bitFieldResult) == 0 || bitFieldResult[0] == 0 {
		return &signCount, nil
	}

	num := bitFieldResult[0]
	for num > 0 {
		if num&1 == 0 {
			break
		}
		signCount++
		num >>= 1
	}

	return &signCount, nil
}
