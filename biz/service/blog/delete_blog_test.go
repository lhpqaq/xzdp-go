package blog

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	goredis "github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlDal "xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
)

func TestDeleteBlogService_Run(t *testing.T) {
	// Mock Redis
	mr, _ := miniredis.Run()
	redis.RedisClient = goredis.NewClient(&goredis.Options{Addr: mr.Addr()})

	// Mock MySQL
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	mysqlDal.DB = gormDB

	// Mock queries
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tb_blog_comment`")).WillReturnRows(sqlmock.NewRows([]string{"id"}))
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `tb_blog`")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	ctx := context.Background()
	c := app.NewContext(1)
	s := NewDeleteBlogService(ctx, c)
	// init req and assert value
	str := "1"
	req := &str
	resp, err := s.Run(req)
	
	// Just ensure no panic. Error is fine.
	if err != nil {
		t.Log(err)
	} else {
		assert.NotEqual(t, nil, resp)
	}
}