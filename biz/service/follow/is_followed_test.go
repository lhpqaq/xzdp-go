package follow

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	follow "xzdp/biz/model/follow"
)

func TestIsFollowedService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewIsFollowedService(ctx, c)
	// init req and assert value
	req := &follow.IsFollowedReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
