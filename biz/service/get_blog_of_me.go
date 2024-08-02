package service

import (
	"context"

	"xzdp/biz/dal/mysql"
	blog "xzdp/biz/model/blog"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type GetBlogOfMeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetBlogOfMeService(Context context.Context, RequestContext *app.RequestContext) *GetBlogOfMeService {
	return &GetBlogOfMeService{RequestContext: RequestContext, Context: Context}
}

func (h *GetBlogOfMeService) Run(req *blog.BlogReq) (resp *[]*blog.Blog, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	blogs, err := mysql.QueryMyBlog(h.Context, 0, utils.GetUser(h.Context))
	if err != nil {
		return nil, err
	}
	return &blogs, nil
}
