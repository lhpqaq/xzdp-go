package blog

import (
	"bytes"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	goredis "github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlDal "xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	"xzdp/biz/model/user"
	"xzdp/biz/utils"
)

func setup() {
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

	// Generic Mocks
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"id"}))
	mock.ExpectExec(".*").WillReturnResult(sqlmock.NewResult(1, 1))
}

func withUser() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u := &user.UserDTO{ID: 1}
		ctx = utils.SaveUser(ctx, u)
		c.Next(ctx)
	}
}

func TestGetHotBlog(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.GET("/blog/hot", GetHotBlog)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/hot", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestGetUserBlog(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.GET("/blog/user/:id", GetUserBlog)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/user/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	if resp.StatusCode() != 200 && resp.StatusCode() != 404 {
		t.Errorf("Unexpected status code: %d", resp.StatusCode())
	}
}

func TestBlogOfMe(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.GET("/blog/of/me", BlogOfMe)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/of/me", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestPostBlog(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.POST("/blog", PostBlog)
	w := ut.PerformRequest(h.Engine, "POST", "/blog", &ut.Body{Body: bytes.NewBufferString("{}"), Len: 2},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestGetBlog(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.GET("/blog/:id", GetBlog)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	if resp.StatusCode() != 200 && resp.StatusCode() != 404 {
		t.Errorf("Unexpected status code: %d", resp.StatusCode())
	}
}

func TestDeleteBlog(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.DELETE("/blog/:id", DeleteBlog)
	w := ut.PerformRequest(h.Engine, "DELETE", "/blog/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestLikeBlog(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.PUT("/blog/like/:id", LikeBlog)
	w := ut.PerformRequest(h.Engine, "PUT", "/blog/like/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestGetLikes(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.GET("/blog/likes/:id", GetLikes)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/likes/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestGetFollowBlog(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	h.GET("/blog/of/follow", GetFollowBlog)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/of/follow", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}
