package blog

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
	"time"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	"xzdp/biz/model/message"
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
	// 获取博客
	var interBlog blog.Blog
	err = mysql.DB.Where("id=?", req).First(&interBlog).Error
	if !errors.Is(err, nil) {
		return nil, errors.New("博客不存在")
	}
	// 判断是否已经点赞
	u := utils.GetUser(h.Context).GetID()
	idStr := strconv.FormatInt(u, 10)
	key := constants.BLOG_LIKED_KEY + *req
	isLike, err := redis.IsLiked(h.Context, key, idStr)
	if err != nil {
		hlog.Debugf("like redis error: %+v", err)
		return nil, err
	}
	fmt.Printf("isLike = %+v", isLike)
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
	// 推送消息
	streamKey := constants.MESSAGE_STREAM_KEY + strconv.FormatInt(interBlog.UserId, 10)
	msg := &message.Message{
		From:    u,
		To:      interBlog.UserId,
		Content: "点赞了你的博客",
		Type:    "like",
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	err = redis.ProduceMq(h.Context, streamKey, msg)
	if err != nil {
		return nil, err
	}
	return &blog.LikeResp{IsLiked: true}, nil
}
