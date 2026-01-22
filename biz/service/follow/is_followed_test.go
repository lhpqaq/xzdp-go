package follow

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"xzdp/biz/model/user"
	"xzdp/biz/utils"
)

func TestIsFollowedService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	
	// Inject User
	u := &user.UserDTO{ID: 1}
	ctx = utils.SaveUser(ctx, u)

	s := NewIsFollowedService(ctx, c)
	// init req and assert value
	req := "1"
	resp, err := s.Run(req)
	
	if err != nil {
		t.Log(err)
	} else {
		assert.NotEqual(t, nil, resp)
	}
}
