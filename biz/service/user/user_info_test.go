package user

import (
	"context"
	"testing"
	model "xzdp/biz/model/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/route/param"
)

// 编写测试用例
func TestUserInfoService_Run(t *testing.T) {
	// 创建一个实例 of UserInfoService, MockDB 和 UserLoginFrom
	h := NewUserInfoService(context.Background(), nil)

	req := &model.UserLoginFrom{
		Phone: "1234567890",
	}

	expectedUser := model.UserInfo{
		UserId: 1,
	}

	mockDB.ExpectQuery("SELECT \\* FROM `tb_user_info` WHERE user_id = \\? ORDER BY `tb_user_info`.`user_id` LIMIT \\?").
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "phone"}).AddRow(1, "1234567890"))

	c := app.RequestContext{
		Params: param.Params{param.Param{"id", "1"}},
	}

	resp, err := h.Run(req, &c)

	assert.DeepEqual(t, &expectedUser, resp)
	assert.DeepEqual(t, nil, err)

	// 如果需要，可以添加更多的测试用例来覆盖不同的场景，例如用户未找到的情况
}
