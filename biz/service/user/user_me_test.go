package user

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	user "xzdp/biz/model/user"
	"xzdp/biz/utils"
)

func TestUserMeService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	
	// Setup user in context
	expectedUser := &user.UserDTO{
		ID: 123,
		NickName: "TestUser",
	}
	ctx = utils.SaveUser(ctx, expectedUser)

	s := NewUserMeService(ctx, c)
	
	req := &user.Empty{}
	resp, err := s.Run(req)
	
	assert.DeepEqual(t, nil, err)
	assert.DeepEqual(t, expectedUser, resp)
}