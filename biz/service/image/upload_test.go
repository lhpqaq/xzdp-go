package image

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	image "xzdp/biz/model/image"
)

func TestUploadService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUploadService(ctx, c)
	// init req and assert value
	req := &[]byte{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
