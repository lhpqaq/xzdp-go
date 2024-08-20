package blog

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	user "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
)

type GetLikesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetLikesService(Context context.Context, RequestContext *app.RequestContext) *GetLikesService {
	return &GetLikesService{RequestContext: RequestContext, Context: Context}
}

func (h *GetLikesService) Run(req *string) (resp *[]*user.UserDTO, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	key := constants.BLOG_LIKED_KEY + *req
	ids, err := redis.RedisClient.ZRange(h.Context, key, 0, 4).Result()
	if err != nil {
		return nil, err
	}
	var users []*user.User
	if !errors.Is(mysql.DB.Where("id in ?", ids).Find(&users).Error, nil) {
		return nil, errors.New("获取失败")
	}
	var userDtos []*user.UserDTO
	for _, u := range users {
		d := &user.UserDTO{
			ID:       u.ID,
			NickName: u.NickName,
			Icon:     u.Icon,
		}
		userDtos = append(userDtos, d)
	}
	if len(userDtos) == 0 {
		userDtos = make([]*user.UserDTO, 0)
	}
	return &userDtos, nil
}
