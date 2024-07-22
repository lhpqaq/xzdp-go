package blog

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	blog "xzdp/biz/model/blog"
	"xzdp/biz/service"
	"xzdp/biz/utils"
)

// GetHotBlog .
// @router /blog/hot [GET]
func GetHotBlog(ctx context.Context, c *app.RequestContext) {
	var err error
	var req blog.BlogReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetHotBlogService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
