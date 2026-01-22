package blog_comment

import (
	"bytes"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	goredis "github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlDal "xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
)

func setup() {
	mr, _ := miniredis.Run()
	redis.RedisClient = goredis.NewClient(&goredis.Options{Addr: mr.Addr()})

	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	mysqlDal.DB = gormDB

	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"id"}))
}

func TestGetHotComment(t *testing.T) {
	setup()
	h := server.Default()
	h.GET("/comment/hot", GetHotComment)
	w := ut.PerformRequest(h.Engine, "GET", "/comment/hot", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestGetComment(t *testing.T) {
	setup()
	h := server.Default()
	h.GET("/comment/:blogID", GetComment)
	w := ut.PerformRequest(h.Engine, "GET", "/comment/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	setup()
	h := server.Default()
	h.POST("/comment/post", PostComment)
	w := ut.PerformRequest(h.Engine, "POST", "/comment/post", &ut.Body{Body: bytes.NewBufferString("{}"), Len: 2},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestLikeComment(t *testing.T) {
	setup()
	h := server.Default()
	h.PUT("/comment/like/:id", LikeComment)
	w := ut.PerformRequest(h.Engine, "PUT", "/comment/like/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestDeleteComment(t *testing.T) {
	setup()
	h := server.Default()
	h.DELETE("/comment/:id", DeleteComment)
	w := ut.PerformRequest(h.Engine, "DELETE", "/comment/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}