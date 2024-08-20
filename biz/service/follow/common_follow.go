package follow

import (
	"context"
	"errors"
	"strconv"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	model "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	follow "xzdp/biz/model/follow"
)

type CommonFollowService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCommonFollowService(Context context.Context, RequestContext *app.RequestContext) *CommonFollowService {
	return &CommonFollowService{RequestContext: RequestContext, Context: Context}
}

func (h *CommonFollowService) Run(targetUserID string) (resp *follow.CommonFollowResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//获取当前用户ID
	user := utils.GetUser(h.Context).GetID()
	key1 := constants.FOLLOW_USER_KEY + strconv.FormatInt(user, 10)
	key2 := constants.FOLLOW_USER_KEY + targetUserID
	arr, err := redis.RedisClient.SInter(h.Context, key1, key2).Result()
	if err != nil {
		return nil, err
	}
	var users []*model.User
	if !errors.Is(mysql.DB.Where("id in ?", arr).Find(&users).Error, nil) {
		return nil, errors.New("查询失败")
	}
	var userDto []*model.UserDTO
	// 遍历arr，转换为userDTO
	for _, u := range users {
		d := utils.UserToUserDTO(u)
		userDto = append(userDto, d)
	}
	return &follow.CommonFollowResp{
		CommonFollows: userDto,
	}, nil
}
