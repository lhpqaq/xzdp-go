package blog_comment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	blog_comment "xzdp/biz/model/blog_comment"
)

type LikeCommentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLikeCommentService(Context context.Context, RequestContext *app.RequestContext) *LikeCommentService {
	return &LikeCommentService{RequestContext: RequestContext, Context: Context}
}

func (h *LikeCommentService) Run(req *string) (resp *blog_comment.LikeResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
