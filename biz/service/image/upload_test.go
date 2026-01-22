package image

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestUploadService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUploadService(ctx, c)
	// init req and assert value
	req := &[]byte{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, err)
	if resp != nil {
		t.Errorf("expected nil response, got %v", resp)
	}
}
