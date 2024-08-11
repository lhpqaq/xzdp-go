package mysql

import (
	"context"
	"xzdp/biz/model/blog"
	"xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func QueryBlogByUserID(ctx context.Context, current int, user *user.UserDTO) (resp []*blog.Blog, err error) {
	var blogs []*blog.Blog
	pageSize := constants.MAX_PAGE_SIZE
	err = DB.Where("user_id = ?", user.ID).Order("liked desc").Limit(pageSize).Offset((current - 1) * pageSize).Find(&blogs).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "query error: %v", err)
		return nil, err
	}
	for i := range blogs {
		blogs[i].NickName = user.NickName
		blogs[i].Icon = user.Icon
	}

	return blogs, nil
}

func QueryHotBlog(ctx context.Context, current int) (resp []*blog.Blog, err error) {
	var blogs []*blog.Blog
	pageSize := constants.MAX_PAGE_SIZE

	if err := DB.Order("liked desc").Limit(pageSize).Offset((current - 1) * pageSize).Find(&blogs).Error; err != nil {
		hlog.CtxErrorf(ctx, "err = %s", err.Error())
		return nil, err
	}

	for i := range blogs {
		u, err := GetById(ctx, blogs[i].UserId)
		if err != nil {
			hlog.CtxErrorf(ctx, "err = %s", err.Error())

			return nil, err
		}
		if err := DB.First(&u, blogs[i].UserId).Error; err != nil {
			hlog.CtxErrorf(ctx, "err = %s", err.Error())

			return nil, err
		}
		blogs[i].NickName = u.NickName
		blogs[i].Icon = u.Icon
	}

	return blogs, nil
}
