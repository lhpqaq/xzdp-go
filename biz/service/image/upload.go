package image

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	image "xzdp/biz/model/image"
)

type UploadService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUploadService(Context context.Context, RequestContext *app.RequestContext) *UploadService {
	return &UploadService{RequestContext: RequestContext, Context: Context}
}

func (h *UploadService) Run(req *[]byte) (resp *image.UploadResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return nil, nil
}
