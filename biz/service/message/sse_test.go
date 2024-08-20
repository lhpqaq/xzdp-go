package message

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	message "xzdp/biz/model/message"
)

func TestSseService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewSseService(ctx, c)
	// init req and assert value
	req := &message.Empty{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
