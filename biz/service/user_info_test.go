package service

import (
	"context"
	"testing"
	model "xzdp/biz/model/user"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// 首先创建一个 Mock 对象来模拟数据库操作
type MockDB struct {
	mock.Mock
}

// 实现一个 MockDB 的方法，以模拟 gorm.DB 的 First 方法
func (mdb *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := mdb.Called(dest, conds)
	if args.Get(0) != nil {
		*dest.(*model.User) = args.Get(0).(model.User)
	}
	return args.Get(1).(*gorm.DB)
}

// 编写测试用例
func TestUserInfoService_Run(t *testing.T) {
	// 创建一个实例 of UserInfoService, MockDB 和 UserLoginFrom
	h := NewUserInfoService(context.Background(), nil)
	mdb := new(MockDB)
	req := &model.UserLoginFrom{
		Phone: "1234567890",
	}
	expectedUser := model.User{
		// 填写期望的用户信息
	}

	// 设置 MockDB 的预期行为和返回结果
	mdb.On("First", mock.AnythingOfType("*model.User"), "phone", req.Phone).Return(expectedUser, mdb)

	// 调用方法并断言预期结果
	resp, err := h.Run(req)
	assert.DeepEqual(t, model.UserResp{}, resp)
	assert.DeepEqual(t, nil, err)

	// 如果需要，可以添加更多的测试用例来覆盖不同的场景，例如用户未找到的情况
}
