package blog

import (
	"context"

	"xzdp/biz/dal/mysql"
	blog "xzdp/biz/model/blog"

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

func (h *GetHotBlogService) Run(req *blog.BlogReq) (resp *[]*blog.Blog, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code

	blogList, err := mysql.QueryHotBlog(h.Context, int(req.Current))
	if err != nil {
		return nil, err
	}
	return &blogList, nil
}
