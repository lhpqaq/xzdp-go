package follow

import (
	"context"
	"errors"
	"strconv"
	follow "xzdp/biz/model/follow"
	service "xzdp/biz/service/follow"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Follow .
// @router /follow [GET]
func Follow(ctx context.Context, c *app.RequestContext) {
	var err error
	id := c.Param("id")
	isFollow := c.Param("isFollow")
	intVal, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	boolVal, err := strconv.ParseBool(isFollow)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	req := follow.FollowReq{
		TargetUser: intVal,
		IsFollow:   boolVal,
	}
	targetUserId := req.GetTargetUser()
	// 如果参数不合法，直接返回
	if targetUserId == 0 {
		utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("参数不合法"))
		return
	}
	resp, err := service.NewFollowService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// IsFollowed .
// @router /follow/isFollowed [GET]
func IsFollowed(ctx context.Context, c *app.RequestContext) {
	var err error
	id := c.Param("id")
	resp, err := service.NewIsFollowedService(ctx, c).Run(id)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CommonFollow .
// @router /follow/common/:id [GET]
func CommonFollow(ctx context.Context, c *app.RequestContext) {
	var err error
	id := c.Param("id")
	resp, err := service.NewCommonFollowService(ctx, c).Run(id)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
