package voucher

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

func TestVoucherList(t *testing.T) {
	setup()
	h := server.Default()
	h.GET("/voucher/list/:id", VoucherList)
	w := ut.PerformRequest(h.Engine, "GET", "/voucher/list/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}

func TestSeckillVoucher(t *testing.T) {
	setup()
	h := server.Default()
	h.POST("/voucher/seckill/:id", SeckillVoucher)
	w := ut.PerformRequest(h.Engine, "POST", "/voucher/seckill/1", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
}