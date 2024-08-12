package blog

import (
	"context"
	"strconv"
	blog "xzdp/biz/model/blog"
	_ "xzdp/biz/model/user"
	"xzdp/biz/service"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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

// PostBlog .
// @router /blog/post [POST]
func PostBlog(ctx context.Context, c *app.RequestContext) {
	var err error
	var req blog.Blog
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewPostBlogService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetBlog .
// @router /blog/:id [GET]
func GetBlog(ctx context.Context, c *app.RequestContext) {
	var err error
	id := c.Param("id")
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetBlogService(ctx, c).Run(&id)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// LikeBlog .
// @router /blog/like/:id [PUT]
func LikeBlog(ctx context.Context, c *app.RequestContext) {
	var err error
	id := c.Param("id")
	resp, err := service.NewLikeBlogService(ctx, c).Run(&id)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetLikes .
// @router /blog/likes/:id [GET]
func GetLikes(ctx context.Context, c *app.RequestContext) {
	var err error
	id := c.Param("id")
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetLikesService(ctx, c).Run(&id)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetUserBlog .
// @router /blog/user/:id [GET]
func GetUserBlog(ctx context.Context, c *app.RequestContext) {
	var err error
	var req blog.BlogReq
	err = c.BindAndValidate(&req)
	id := c.Param("id")
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewGetUserBlogService(ctx, c).Run(&req, userId)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteBlog .
// @router /blog/:id [DELETE]
func DeleteBlog(ctx context.Context, c *app.RequestContext) {
	var err error
	req := c.Param("id")
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewDeleteBlogService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetFollowBlog .
// @router /blog/of/follow [GET]
func GetFollowBlog(ctx context.Context, c *app.RequestContext) {
	var err error
	var req blog.FollowBlogReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetFollowBlogService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// BlogOfMe .
// @router /blog/of/me [GET]
func BlogOfMe(ctx context.Context, c *app.RequestContext) {
	var err error
	var req blog.BlogReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewBlogOfMeService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
