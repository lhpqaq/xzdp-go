package user

import (
	"time"

	"gorm.io/gorm"
)

func (User) TableName() string {
	return "tb_user"
}

func (UserInfo) TableName() string {
	return "tb_user_info"
}

// BeforeCreate 是一个 GORM 钩子函数，在插入数据前执行
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	u.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return
}

func (u *UserInfo) BeforeCreate(tx *gorm.DB) (err error) {
	u.Fans = 0
	u.Birthday = time.Now().Format("2006-01-02 15:04:05")
	u.Followee = 0
	u.Credits = 0
	u.Gender = 1
	u.Introduce = ""
	u.Level = 0
	u.City = "北京"
	u.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	u.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return
}

// func (u UserDTO) MarshalBinary() ([]byte, error) {
// 	return json.Marshal(u)
// }
