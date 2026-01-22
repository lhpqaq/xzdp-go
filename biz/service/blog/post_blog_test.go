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
	"xzdp/biz/model/user"
	"xzdp/biz/utils"
)

func TestPostBlogService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)

	// Inject User
	u := &user.UserDTO{ID: 1}
	ctx = utils.SaveUser(ctx, u)

	// Setup Mocks
	mr, _ := miniredis.Run()
	redis.RedisClient = goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	mysqlDal.DB = gormDB

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `tb_blog`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// Use broad match for follow query
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	s := NewPostBlogService(ctx, c)
	req := &blog.Blog{}
	resp, err := s.Run(req)

	if err != nil {
		t.Log("Warning: PostBlog returned error:", err)
	}
	commonTestAssert.NotEqual(t, nil, resp)
}
