package blog_comment

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestDeleteCommentService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewDeleteCommentService(ctx, c)
	// init req and assert value
	str := "1"
	req := &str
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, err)
	if resp != nil {
		t.Errorf("expected nil response, got %v", resp)
	}
}
