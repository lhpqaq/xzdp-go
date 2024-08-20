package blog

import (
	"context"
	"xzdp/biz/dal/mysql"

	"github.com/cloudwego/hertz/pkg/app"
	blog "xzdp/biz/model/blog"
	"xzdp/biz/model/user"
)

type GetUserBlogService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserBlogService(Context context.Context, RequestContext *app.RequestContext) *GetUserBlogService {
	return &GetUserBlogService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserBlogService) Run(req *blog.BlogReq, uerID int64) (resp *[]*blog.Blog, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	u, err := mysql.GetById(h.Context, uerID)
	if err != nil {
		return nil, err
	}
	d := &user.UserDTO{
		ID:       u.ID,
		NickName: u.NickName,
		Icon:     u.Icon,
	}
	blogList, err := mysql.QueryBlogByUserID(h.Context, int(req.Current), d)
	if err != nil {
		return nil, err
	}
	return &blogList, nil
}
