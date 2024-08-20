package follow

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	follow "xzdp/biz/model/follow"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"
)

type FollowService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewFollowService(Context context.Context, RequestContext *app.RequestContext) *FollowService {
	return &FollowService{RequestContext: RequestContext, Context: Context}
}

func (h *FollowService) Run(req *follow.FollowReq) (resp *follow.FollowResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	myID := utils.GetUser(h.Context).GetID()
	isFollow := req.GetIsFollow()
	targetUserId := req.GetTargetUser()
	f := follow.Follow{
		UserId:       myID,
		FollowUserId: targetUserId,
	}
	// 如果是true,则添加关注，将用户id和被关注用户的id存入数据库
	if isFollow {
		// 判断是否已经关注
		if !errors.Is(redis.RedisClient.SIsMember(h.Context, constants.FOLLOW_USER_KEY+strconv.FormatInt(myID, 10), targetUserId).Err(), nil) {
			return nil, errors.New("关注失败")
		}
		// 将关注的用户存入redis的set中
		if !errors.Is(redis.RedisClient.SAdd(h.Context, constants.FOLLOW_USER_KEY+strconv.FormatInt(myID, 10), targetUserId).Err(), nil) {
			hlog.CtxErrorf(h.Context, "err = %v", err)
			return nil, err
		}
		if !errors.Is(mysql.DB.Create(&f).Error, nil) {
			return nil, errors.New("关注失败")
		}
		return &follow.FollowResp{RespBody: &f}, nil
	}
	// 如果是false,则取消关注
	if !errors.Is(mysql.DB.Where("user_id = ? and follow_user_id = ?", myID, targetUserId).Delete(&f).Error, nil) {
		return nil, errors.New("取消关注失败")
	}
	// 将取消关注的用户从redis的set中删除
	if !errors.Is(redis.RedisClient.SRem(h.Context, constants.FOLLOW_USER_KEY+strconv.FormatInt(myID, 10), targetUserId).Err(), nil) {
		hlog.CtxErrorf(h.Context, "err = %v", err)
		return nil, err
	}
	return &follow.FollowResp{RespBody: &f}, nil
}
