package user

import (
	"context"
	"strconv"

	"xzdp/biz/dal/mysql"
	user "xzdp/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type UserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UserInfoService {
	return &UserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UserInfoService) Run(req *user.UserLoginFrom, c *app.RequestContext) (resp *user.UserInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return nil, err
	}
	userInfo, err := mysql.GetUserInfoById(h.Context, id)
	if err == nil && userInfo != nil {
		return userInfo, nil
	}

	err = h.createNewUserWithId(id)
	if err != nil {
		return nil, err
	}
	userInfo, err = mysql.GetUserInfoById(h.Context, id)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (h *UserInfoService) createNewUserWithId(id int64) error {
	user := user.UserInfo{
		UserId: id,
	}
	result := mysql.DB.Create(&user)
	hlog.CtxDebugf(h.Context, "result = %+v", result)
	return result.Error
}
