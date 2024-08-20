package blog_comment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	blog_comment "xzdp/biz/model/blog_comment"
)

type PostCommentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPostCommentService(Context context.Context, RequestContext *app.RequestContext) *PostCommentService {
	return &PostCommentService{RequestContext: RequestContext, Context: Context}
}

func (h *PostCommentService) Run(req *blog_comment.BlogComment) (resp *blog_comment.BlogComment, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
