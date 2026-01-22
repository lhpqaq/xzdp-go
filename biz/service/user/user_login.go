package user

import (
	"context"
	"errors"
	"fmt"

	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	model "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	red "github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DB             *gorm.DB
	Redis          *red.Client
}

func NewUserLoginService(ctx context.Context, requestContext *app.RequestContext) *UserLoginService {
	return &UserLoginService{
		RequestContext: requestContext,
		Context:        ctx,
		DB:             mysql.DB,
		Redis:          redis.RedisClient,
	}
}

// NewUserLoginServiceWithDB is used for testing or when custom DB/Redis are needed
func NewUserLoginServiceWithDB(ctx context.Context, requestContext *app.RequestContext, db *gorm.DB, rds *red.Client) *UserLoginService {
	return &UserLoginService{
		RequestContext: requestContext,
		Context:        ctx,
		DB:             db,
		Redis:          rds,
	}
}

func (h *UserLoginService) Run(req *model.UserLoginFrom) (resp *model.Result, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	phone := req.Phone
	code := req.Code
	if phone == "" || code == "" {
		return nil, errors.New("phone or code can't be empty")
	}
	redisCode, err := h.Redis.Get(h.Context, constants.LOGIN_CODE_KEY+phone).Result()
	if err != nil {
		if errors.Is(err, red.Nil) {
			return nil, fmt.Errorf("code expired or not found")
		}
		hlog.CtxErrorf(h.Context, "err = %s", err.Error())
		return nil, err
	}
	if redisCode != code {
		return nil, fmt.Errorf("code not match")
	}

	token, err := utils.RandomUUID()
	if err != nil {
		return nil, err
	}

	var user model.User
	result := h.DB.Debug().First(&user, "phone = ?", phone)
	hlog.CtxInfof(h.Context, "result = %+v", result)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			user, err = h.createNewUserWithPhone(phone)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, result.Error
		}
	}

	var userdto model.UserDTO
	copier.Copy(&userdto, &user)
	if err = h.Redis.HMSet(h.Context, constants.LOGIN_USER_KEY+token, map[string]interface{}{
		"id":        userdto.ID,
		"nick_name": userdto.NickName,
		"icon":      userdto.Icon,
	}).Err(); err != nil {
		hlog.CtxErrorf(h.Context, "err = %s", err.Error())
		hlog.CtxErrorf(h.Context, "userdto = %+v", userdto)
		return nil, err
	}

	return &model.Result{Success: true, Data: &token}, nil
}

func (h *UserLoginService) createNewUserWithPhone(phone string) (model.User, error) {
	user := model.User{
		Phone:    phone,
		NickName: "user_" + utils.RandomString(10),
	}

	result := h.DB.Debug().Create(&user)
	hlog.CtxInfof(h.Context, "result = %+v", result)
	return user, result.Error
}
