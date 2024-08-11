package follow

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	follow "xzdp/biz/model/follow"
	"xzdp/biz/service"
	"xzdp/biz/utils"
)

// Follow .
// @router /follow [GET]
func Follow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req follow.FollowReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
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
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCommonFollowService(ctx, c).Run(id)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
