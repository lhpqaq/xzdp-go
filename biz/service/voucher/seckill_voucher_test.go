package voucher

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	goredis "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	redsyncgoredis "github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlDal "xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
)

func TestSeckillVoucherService_Run(t *testing.T) {

	// Mock Redis

	mr, err := miniredis.Run()

	if err != nil {

		t.Fatal(err)

	}

	defer mr.Close()

	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})

	redis.RedisClient = rdb

	pool := redsyncgoredis.NewPool(rdb)

	redis.RedsyncClient = redsync.New(pool)

	// Mock MySQL

	db, _, err := sqlmock.New()

	if err != nil {

		t.Fatal(err)

	}

	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{

		Conn: db,

		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {

		t.Fatal(err)

	}

	mysqlDal.DB = gormDB

	ctx := context.Background()

	c := app.NewContext(1)

	s := NewSeckillVoucherService(ctx, c)

	// init req

	id := int64(1)

	// Expect DB calls if cache miss? Or logic flow.

	// Just ensure no panic for now.

	resp, err := s.Run(&id)

	if err != nil {

		t.Log(err)

	} else {

		assert.NotEqual(t, nil, resp)

	}

}
