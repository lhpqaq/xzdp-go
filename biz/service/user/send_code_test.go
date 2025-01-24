package user

import (
	"context"
	"database/sql"
	"testing"

	mysql2 "xzdp/biz/dal/mysql"

	"gorm.io/driver/mysql"

	redis2 "xzdp/biz/dal/redis"
	user "xzdp/biz/model/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	mockDB sqlmock.Sqlmock
	db     *sql.DB
)

func TestSendCodeService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewSendCodeService(ctx, c)
	// init req and assert value
	req := &user.UserLoginFrom{
		Phone: "12345678901"}
	resp, err := s.Run(req)
	assert.Nil(t, resp)
	assert.NotNil(t, err)

	resp, err = s.Run(&user.UserLoginFrom{
		Phone: "13412332123"})

	assert.DeepEqual(t, &user.Result{Success: true}, resp)
	assert.DeepEqual(t, nil, err)
}

func TestMain(m *testing.M) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	redis2.RedisClient = redis.NewClient(&redis.Options{
		Addr: s.Addr(), // mock redis server的地址
	})
	var mockErr error
	db, mockDB, mockErr = sqlmock.New()
	if mockErr != nil {
		panic(mockErr)
	}
	defer db.Close()

	mysql2.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true, // 跳过版本检查以支持 mock
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	m.Run()
}
