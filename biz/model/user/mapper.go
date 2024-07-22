package user

import (
	"time"

	"gorm.io/gorm"
)

func (User) TableName() string {
	return "tb_user"
}

// BeforeCreate 是一个 GORM 钩子函数，在插入数据前执行
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	u.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return
}

// func (u UserDTO) MarshalBinary() ([]byte, error) {
// 	return json.Marshal(u)
// }
