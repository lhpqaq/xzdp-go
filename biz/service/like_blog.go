package service

import (
	"context"
	"errors"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
	"time"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	blog "xzdp/biz/model/blog"
)

type LikeBlogService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLikeBlogService(Context context.Context, RequestContext *app.RequestContext) *LikeBlogService {
	return &LikeBlogService{RequestContext: RequestContext, Context: Context}
}

func (h *LikeBlogService) Run(req *string) (resp *blog.LikeResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// 判断是否已经点赞
	u := utils.GetUser(h.Context).GetID()
	idStr := strconv.FormatInt(u, 10)
	key := constants.BLOG_LIKED_KEY + *req
	isLike, err := redis.IsLiked(h.Context, key, idStr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("score = %+v", isLike)
	// 如果已经点赞则取消点赞
	if isLike {
		if !errors.Is(redis.RedisClient.ZRem(h.Context, key, idStr).Err(), nil) {
			return nil, errors.New("取消点赞失败")
		}
		// 同步减少点赞数
		mysql.DB.Model(&blog.Blog{}).Where("id = ?", req).UpdateColumn("liked", gorm.Expr("liked - ?", 1))
		return &blog.LikeResp{IsLiked: false}, nil
	}
	// 否则点赞
	if !errors.Is(redis.RedisClient.ZAdd(h.Context, key, &redis2.Z{
		Score:  float64(time.Now().Unix()),
		Member: idStr,
	}).Err(), nil) {
		return nil, errors.New("点赞失败")
	}
	// 同步增加点赞数
	mysql.DB.Model(&blog.Blog{}).Where("id = ?", req).UpdateColumn("liked", gorm.Expr("liked + ?", 1))
	return &blog.LikeResp{IsLiked: true}, nil
}
