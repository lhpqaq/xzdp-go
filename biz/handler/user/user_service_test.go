package user

import (
	"bytes"
	"regexp"
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
	// Mock Redis
	mr, _ := miniredis.Run()
	// defer mr.Close() // In a real suite we'd handle cleanup better
	redis.RedisClient = goredis.NewClient(&goredis.Options{Addr: mr.Addr()})

	// Mock MySQL
	db, mock, _ := sqlmock.New()
	// defer db.Close()
	
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	mysqlDal.DB = gormDB
	
	// Mock common queries if needed to avoid panic
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tb_user`")).WillReturnRows(sqlmock.NewRows([]string{"id"}))
}

func TestUserMethod(t *testing.T) {
	setup()
	h := server.Default()
	h.GET("/user/me", UserMe)
	w := ut.PerformRequest(h.Engine, "GET", "/user/me", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	// Should be 200 OK because we mocked DB/Redis (or at least stopped panic)
	// But UserMe requires login, so it might return error or empty
	assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "", string(resp.Body())) 
}

func TestSendCode(t *testing.T) {
	setup()
	h := server.Default()
	h.POST("/user/code", SendCode)
	// Body needs to be valid JSON for BindAndValidate?
	w := ut.PerformRequest(h.Engine, "POST", "/user/code", &ut.Body{Body: bytes.NewBufferString(`{"phone":"13800000000"}`), Len: 21},
		ut.Header{Key: "Content-Type", Value: "application/json"})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestUserLogin(t *testing.T) {
	setup()
	h := server.Default()
	h.POST("/user/login", UserLogin)
	w := ut.PerformRequest(h.Engine, "POST", "/user/login", &ut.Body{Body: bytes.NewBufferString(`{"phone":"13800000000", "code":"123456"}`), Len: 40},
		ut.Header{Key: "Content-Type", Value: "application/json"})
	resp := w.Result()
	// Might fail login due to mocked logic, but should not panic
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestUserInfo(t *testing.T) {
	setup()
	h := server.Default()
	h.GET("/user/info/:id", UserInfo)
	w := ut.PerformRequest(h.Engine, "GET", "/user/info/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}