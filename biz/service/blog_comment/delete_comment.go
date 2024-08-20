package blog_comment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	blog_comment "xzdp/biz/model/blog_comment"
)

type DeleteCommentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteCommentService(Context context.Context, RequestContext *app.RequestContext) *DeleteCommentService {
	return &DeleteCommentService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteCommentService) Run(req *string) (resp *blog_comment.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
