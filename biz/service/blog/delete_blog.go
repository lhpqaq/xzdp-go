package blog

import (
	"context"
	"errors"
	"strconv"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
	blog "xzdp/biz/model/blog"
)

type DeleteBlogService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteBlogService(Context context.Context, RequestContext *app.RequestContext) *DeleteBlogService {
	return &DeleteBlogService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteBlogService) Run(req *string) (resp *blog.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	key := constants.BLOG_LIKED_KEY + *req
	// 从redis删除点赞数据
	ok, err := redis.HasLikes(h.Context, key)
	if err != nil {
		return nil, err
	}
	if ok {
		if !errors.Is(redis.DeleteLikes(h.Context, key), nil) {
			return nil, err
		}
	}
	// 删除评论信息
	bid, err := strconv.ParseInt(*req, 10, 64)
	if err != nil {
		return nil, err
	}
	err = mysql.DeleteBlogComment(h.Context, bid)
	if !errors.Is(err, nil) {
		return nil, err
	}
	// 从数据库删除博客
	err = mysql.DB.Where("id = ?", req).Delete(&blog.Blog{}).Error
	if !errors.Is(err, nil) {
		return nil, err
	}
	return &blog.Empty{}, nil
}
