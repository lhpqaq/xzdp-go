package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	follow "xzdp/biz/model/follow"
	user "xzdp/biz/model/user"
)

func GetFansByID(ctx context.Context, id int64) (resp []*user.UserDTO, err error) {
	var fs []follow.Follow
	err = DB.Where("follow_user_id = ?", id).Find(&fs).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return make([]*user.UserDTO, 0), nil
	}
	if err != nil {
		return nil, errors.New("获取失败")
	}
	for _, f := range fs {
		interUser, e := GetById(ctx, f.UserId)
		if e != nil {
			return nil, e
		}
		u := &user.UserDTO{
			ID:       interUser.ID,
			NickName: interUser.NickName,
			Icon:     interUser.Icon,
		}
		resp = append(resp, u)
	}
	return resp, nil
}
