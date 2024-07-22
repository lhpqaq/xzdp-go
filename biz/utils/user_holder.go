package utils

import (
	"context"
	model "xzdp/biz/model/user"
)

func SaveUser(ctx context.Context, user *model.UserDTO) context.Context {
	return context.WithValue(ctx, "user", user)
}

// GetUser retrieves the user information from the context.
func GetUser(ctx context.Context) *model.UserDTO {
	user, ok := ctx.Value("user").(*model.UserDTO)
	if !ok {
		return nil
	}
	return user
}

// RemoveUser creates a new context without the user information.
func RemoveUser(ctx context.Context) context.Context {
	return context.WithValue(ctx, "user", nil)
}
