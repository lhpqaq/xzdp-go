package user

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
	model "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
)

func TestUserLoginService_Run(t *testing.T) {
	// Setup miniredis
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mr.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	// Setup sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database connection", err)
	}

	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUserLoginServiceWithDB(ctx, c, gormDB, rdb)

	phone := "13800000000"
	code := "123456"

	// Case 1: Success (User exists)
	t.Run("Success_UserExists", func(t *testing.T) {
		// Mock Redis
		mr.Set(constants.LOGIN_CODE_KEY+phone, code)

		// Mock DB
		rows := sqlmock.NewRows([]string{"id", "phone", "nick_name", "icon"}).
			AddRow(1, phone, "test_user", "icon.png")
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tb_user` WHERE phone = ? ORDER BY `tb_user`.`id` LIMIT ?")).
			WithArgs(phone, 1).
			WillReturnRows(rows)

		req := &model.UserLoginFrom{
			Phone: phone,
			Code:  code,
		}
		resp, err := s.Run(req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, true, resp.Success)
		assert.NotEqual(t, nil, resp.Data)
	})

	// Case 2: Success (User does not exist, create new)
	t.Run("Success_UserNew", func(t *testing.T) {
		mr.Set(constants.LOGIN_CODE_KEY+phone, code)

		// Mock DB - First query returns not found
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tb_user` WHERE phone = ? ORDER BY `tb_user`.`id` LIMIT ?")).
			WithArgs(phone, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		// Mock DB - Create user
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `tb_user`")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()). // Adjust args based on model
			WillReturnResult(sqlmock.NewResult(2, 1))
		mock.ExpectCommit()

		req := &model.UserLoginFrom{
			Phone: phone,
			Code:  code,
		}
		resp, err := s.Run(req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, true, resp.Success)
	})

	// Case 3: Code mismatch
	t.Run("Fail_CodeMismatch", func(t *testing.T) {
		mr.Set(constants.LOGIN_CODE_KEY+phone, "654321")

		req := &model.UserLoginFrom{
			Phone: phone,
			Code:  code,
		}
		resp, err := s.Run(req)
		assert.NotEqual(t, nil, err)
		assert.DeepEqual(t, "code not match", err.Error())
		assert.DeepEqual(t, (*model.Result)(nil), resp)
	})

	// Case 4: Code expired/not found
	t.Run("Fail_CodeExpired", func(t *testing.T) {
		mr.FlushAll()

		req := &model.UserLoginFrom{
			Phone: phone,
			Code:  code,
		}
		resp, err := s.Run(req)
		assert.NotEqual(t, nil, err)
		assert.DeepEqual(t, "code expired or not found", err.Error())
		assert.DeepEqual(t, (*model.Result)(nil), resp)
	})
}