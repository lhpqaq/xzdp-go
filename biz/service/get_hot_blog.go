package service

import (
	"context"

	"xzdp/biz/dal/mysql"
	blog "xzdp/biz/model/blog"
	"xzdp/biz/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type GetHotBlogService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetHotBlogService(Context context.Context, RequestContext *app.RequestContext) *GetHotBlogService {
	return &GetHotBlogService{RequestContext: RequestContext, Context: Context}
}

func QueryHotBlog(ctx context.Context, current int) (resp []*blog.Blog, err error) {
	var blogs []*blog.Blog
	pageSize := constants.MAX_PAGE_SIZE

	if err := mysql.DB.Order("liked desc").Limit(pageSize).Offset((current - 1) * pageSize).Find(&blogs).Error; err != nil {
		hlog.CtxErrorf(ctx, "err = %s", err.Error())
		return nil, err
	}

	for i := range blogs {
		user, err := mysql.GetById(ctx, blogs[i].UserId)
		if err != nil {
			hlog.CtxErrorf(ctx, "err = %s", err.Error())

			return nil, err
		}
		if err := mysql.DB.First(&user, blogs[i].UserId).Error; err != nil {
			hlog.CtxErrorf(ctx, "err = %s", err.Error())

			return nil, err
		}
		blogs[i].Name = user.NickName
		blogs[i].Icon = user.Icon
	}

	return blogs, nil
}

func (h *GetHotBlogService) Run(req *blog.BlogReq) (resp *[]*blog.Blog, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code

	blogList, err := QueryHotBlog(h.Context, int(req.Current))
	if err != nil {
		return nil, err
	}
	return &blogList, nil
}
