package blog

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	blog "xzdp/biz/model/blog"
)

func TestBlogOfMeService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewBlogOfMeService(ctx, c)
	// init req and assert value
	req := &blog.BlogReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
