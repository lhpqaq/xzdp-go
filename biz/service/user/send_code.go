package user

import (
	"context"
	"errors"
	"time"
	"xzdp/biz/dal/redis"
	user "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type SendCodeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSendCodeService(Context context.Context, RequestContext *app.RequestContext) *SendCodeService {
	return &SendCodeService{RequestContext: RequestContext, Context: Context}
}

func (h *SendCodeService) Run(req *user.UserLoginFrom) (resp *user.Result, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	phone := req.Phone

	if !utils.ValidateMobile(phone) {
		return nil, errors.New("phone isn't validate")
	}
	if phone == "" {
		return nil, errors.New("phone can't be empty")
	}

	code := utils.GenerateDigits(6)
	err = redis.RedisClient.Set(h.Context, constants.LOGIN_CODE_KEY+phone, code, constants.LOGIN_CODE_EXPIRE*time.Second).Err() // add expiration
	if err != nil {
		hlog.CtxErrorf(h.Context, "err = %s", err.Error())
		return nil, err
	}

	hlog.CtxInfof(h.Context, "code = %s", code)
	return &user.Result{Success: true}, nil
}
