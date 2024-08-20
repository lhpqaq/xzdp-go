package blog

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	blog "xzdp/biz/model/blog"
)

func TestLikeBlogService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewLikeBlogService(ctx, c)
	// init req and assert value
	req := &string{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
