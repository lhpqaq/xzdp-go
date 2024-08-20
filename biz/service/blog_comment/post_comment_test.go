package blog_comment

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	blog_comment "xzdp/biz/model/blog_comment"
)

func TestPostCommentService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewPostCommentService(ctx, c)
	// init req and assert value
	req := &blog_comment.BlogComment{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
