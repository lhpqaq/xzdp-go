package follow

import (
	"context"
	"errors"
	"strconv"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	follow "xzdp/biz/model/follow"
)

type IsFollowedService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewIsFollowedService(Context context.Context, RequestContext *app.RequestContext) *IsFollowedService {
	return &IsFollowedService{RequestContext: RequestContext, Context: Context}
}

func (h *IsFollowedService) Run(targetUserID string) (resp *follow.IsFollowedResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//获取当前用户ID
	user := utils.GetUser(h.Context).GetID()
	//查找是否关注
	if !errors.Is(redis.RedisClient.SIsMember(h.Context, constants.FOLLOW_USER_KEY+strconv.FormatInt(user, 10), targetUserID).Err(), nil) {
		return &follow.IsFollowedResp{IsFollowed: false}, nil
	}
	return &follow.IsFollowedResp{IsFollowed: true}, nil
}
