package user

import (
	"context"

	user "xzdp/biz/model/user"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type UserMeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserMeService(Context context.Context, RequestContext *app.RequestContext) *UserMeService {
	return &UserMeService{RequestContext: RequestContext, Context: Context}
}

func (h *UserMeService) Run(req *user.Empty) (resp *user.UserDTO, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	userdto := utils.GetUser(h.Context)
	return userdto, nil
}
