package message

import (
	"bytes"
	"context"
	"testing"
	"time"

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

func withUser() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u := &user.UserDTO{ID: 1}
		ctx = utils.SaveUser(ctx, u)
		c.Next(ctx)
	}
}

func TestSse(t *testing.T) {
	setup()
	h := server.Default()
	h.Use(withUser())
	// Wrap Sse to inject a timeout context to break the infinite loop
	h.GET("/message/sse", func(ctx context.Context, c *app.RequestContext) {
		timeoutCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()
		Sse(timeoutCtx, c)
	})
	
	w := ut.PerformRequest(h.Engine, "GET", "/message/sse", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}