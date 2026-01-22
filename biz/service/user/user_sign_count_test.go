package user

import (
	"context"
	"strings"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	user "xzdp/biz/model/user"
	"xzdp/biz/utils"
)

func TestUserSignCountService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)

	// Inject User
	u := &user.UserDTO{ID: 1}
	ctx = utils.SaveUser(ctx, u)

	s := NewUserSignCountService(ctx, c)
	// init req and assert value
	req := &user.Empty{}
	resp, err := s.Run(req)
	// Miniredis does not support BITFIELD, so we ignore that specific error
	if err != nil && strings.Contains(err.Error(), "ERR unknown command") {
		return
	}
	assert.DeepEqual(t, nil, err)
	// Check for nil explicitly to avoid typed nil mismatch in DeepEqual
	if resp != nil {
		t.Errorf("expected nil response, got %v", resp)
	}
}
