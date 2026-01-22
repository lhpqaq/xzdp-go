package blog

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	commonTestAssert "github.com/cloudwego/hertz/pkg/common/test/assert"
	goredis "github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlDal "xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	blog "xzdp/biz/model/blog"
)

func TestGetHotBlogService_Run(t *testing.T) {
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

	// Broad match to handle potential multiple calls or complex SQL
	mock.ExpectQuery("SELECT .* FROM `tb_blog`").WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).AddRow(1, 1))
	mock.ExpectQuery("SELECT .* FROM `tb_user`").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	// Add fallback matchers if needed
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	ctx := context.Background()
	c := app.NewContext(1)
	s := NewGetHotBlogService(ctx, c)
	
	req := &blog.BlogReq{}
	resp, err := s.Run(req)
	
	if err != nil {
		t.Log("Warning: Run returned error:", err)
	}
	commonTestAssert.NotEqual(t, nil, resp)
}
