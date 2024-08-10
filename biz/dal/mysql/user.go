package mysql

import (
	"context"
	"errors"
	model "xzdp/biz/model/user"
	"xzdp/biz/utils"

	"gorm.io/gorm"
)

func GetById(ctx context.Context, id int64) (*model.User, error) {

	var user model.User
	db := DB.Model(&model.User{})

	// Perform the query
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle case where no user is found
			return nil, utils.ErrNotFound
		}
		// Handle other potential errors
		return nil, err
	}

	// Return the user
	return &user, nil
}

func GetUserInfoById(ctx context.Context, id int64) (*model.UserInfo, error) {

	var user model.UserInfo
	db := DB.Model(&model.UserInfo{})

	// Perform the query
	if err := db.First(&user, "user_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle case where no user is found
			return nil, utils.ErrNotFound
		}
		// Handle other potential errors
		return nil, err
	}

	// Return the user
	return &user, nil
}
