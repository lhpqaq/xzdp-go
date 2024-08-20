package blog_comment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	blog_comment "xzdp/biz/model/blog_comment"
)

type GetHotCommentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetHotCommentService(Context context.Context, RequestContext *app.RequestContext) *GetHotCommentService {
	return &GetHotCommentService{RequestContext: RequestContext, Context: Context}
}

func (h *GetHotCommentService) Run(req *blog_comment.CommentReq) (resp *[]*blog_comment.BlogComment, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
