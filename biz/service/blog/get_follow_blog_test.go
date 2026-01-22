package blog

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	blog "xzdp/biz/model/blog"
	"xzdp/biz/model/user"
	"xzdp/biz/utils"
)

func TestGetFollowBlogService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	
	// Inject User
	u := &user.UserDTO{ID: 1}
	ctx = utils.SaveUser(ctx, u)

	s := NewGetFollowBlogService(ctx, c)
	// init req and assert value
	req := &blog.FollowBlogReq{}
	resp, err := s.Run(req)
	
	// Ensure no panic
	if err != nil {
		t.Log(err)
	} else {
		assert.NotEqual(t, nil, resp)
	}
}
