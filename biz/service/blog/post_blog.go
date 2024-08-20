package blog

import (
	"context"
	"errors"
	redis2 "github.com/go-redis/redis/v8"
	"strconv"
	"time"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	blog "xzdp/biz/model/blog"
)

type PostBlogService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPostBlogService(Context context.Context, RequestContext *app.RequestContext) *PostBlogService {
	return &PostBlogService{RequestContext: RequestContext, Context: Context}
}

func (h *PostBlogService) Run(req *blog.Blog) (resp *blog.Blog, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	u := utils.GetUser(h.Context).GetID()
	req.UserId = u
	if !errors.Is(mysql.DB.Create(&req).Error, nil) {
		return nil, errors.New("创建失败")
	}
	req.Icon = utils.GetUser(h.Context).GetIcon()
	req.NickName = utils.GetUser(h.Context).GetNickName()
	req.IsLiked = false
	fans, err := mysql.GetFansByID(h.Context, u)
	if err != nil {
		return nil, err
	}
	for _, fan := range fans {
		key := constants.FEED_KEY + strconv.FormatInt(fan.ID, 10)
		err = redis.RedisClient.ZAdd(h.Context, key, &redis2.Z{
			Score:  float64(time.Now().Unix()),
			Member: req.ID,
		}).Err()
	}
	return req, nil
}
