package blog

import (
	"context"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/model/user"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	blog "xzdp/biz/model/blog"
)

type BlogOfMeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBlogOfMeService(Context context.Context, RequestContext *app.RequestContext) *BlogOfMeService {
	return &BlogOfMeService{RequestContext: RequestContext, Context: Context}
}

func (h *BlogOfMeService) Run(req *blog.BlogReq) (resp *[]*blog.Blog, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	u := utils.GetUser(h.Context)
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
