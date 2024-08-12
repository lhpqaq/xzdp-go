package utils

import "xzdp/biz/model/user"

func UserToUserDTO(u *user.User) *user.UserDTO {
	return &user.UserDTO{
		ID:       u.ID,
		NickName: u.NickName,
		Icon:     u.Icon,
	}
}
