package follow

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	follow "xzdp/biz/model/follow"
)

func TestFollowService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewFollowService(ctx, c)
	// init req and assert value
	req := &follow.FollowReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
