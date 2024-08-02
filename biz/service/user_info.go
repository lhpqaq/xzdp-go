package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"xzdp/biz/dal/mysql"

	model "xzdp/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UserInfoService {
	return &UserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UserInfoService) Run(id string) (resp *model.UserResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", id)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	var user model.User
	if errors.Is(mysql.DB.First(&user, "id = ?", id).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	// 将用户信息返回
	resp.RespBody = &user
	fmt.Println(user)
	return
}
