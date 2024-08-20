package blog

import (
	"context"
	"testing"

	blog "xzdp/biz/model/blog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestGetHotBlogService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetHotBlogService(ctx, c)
	// init req and assert value
	req := &blog.BlogReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
