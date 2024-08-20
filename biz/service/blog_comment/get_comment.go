package blog_comment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	blog_comment "xzdp/biz/model/blog_comment"
)

type GetCommentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCommentService(Context context.Context, RequestContext *app.RequestContext) *GetCommentService {
	return &GetCommentService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCommentService) Run(req *blog_comment.CommentReq) (resp *[]*blog_comment.BlogComment, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
