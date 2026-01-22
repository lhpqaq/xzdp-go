package user

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/route/param"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "xzdp/biz/model/user"
)

func TestUserInfoService_Run(t *testing.T) {
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
	c.Params = param.Params{
		{Key: "id", Value: "1"},
	}

	s := NewUserInfoServiceWithDB(ctx, c, gormDB)

	// Case 1: Success (User Info exists)
	t.Run("Success_UserExists", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id", "introduce"}).
			AddRow(1, "hello world")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tb_user_info` WHERE user_id = ? ORDER BY `tb_user_info`.`user_id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnRows(rows)

		req := &model.UserLoginFrom{}
		resp, err := s.Run(req, c)

		assert.DeepEqual(t, nil, err)
		assert.NotEqual(t, nil, resp)
		assert.DeepEqual(t, int64(1), resp.UserId)
		assert.DeepEqual(t, "hello world", resp.Introduce)
	})

	// Case 2: Success (User Info creates)
	t.Run("Success_UserCreate", func(t *testing.T) {
		// 1. Get fails
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tb_user_info` WHERE user_id = ? ORDER BY `tb_user_info`.`user_id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		// 2. Create
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `tb_user_info`")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()). // 11 fields
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// 3. Get success
		rows := sqlmock.NewRows([]string{"user_id"}).AddRow(1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tb_user_info` WHERE user_id = ? ORDER BY `tb_user_info`.`user_id` LIMIT ?")).
			WithArgs(1, 1).
			WillReturnRows(rows)

		req := &model.UserLoginFrom{}
		resp, err := s.Run(req, c)

		assert.DeepEqual(t, nil, err)
		assert.NotEqual(t, nil, resp)
		assert.DeepEqual(t, int64(1), resp.UserId)
	})
}
