package follow

import (
	"context"
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
	follow "xzdp/biz/model/follow"
	"xzdp/biz/model/user"
	"xzdp/biz/utils"
)

func TestFollowService_Run(t *testing.T) {
	// Mock Redis
	mr, _ := miniredis.Run()
	redis.RedisClient = goredis.NewClient(&goredis.Options{Addr: mr.Addr()})

	// Mock MySQL
	db, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	mysqlDal.DB = gormDB

	ctx := context.Background()
	c := app.NewContext(1)

	// Inject User
	u := &user.UserDTO{ID: 1}
	ctx = utils.SaveUser(ctx, u)

	s := NewFollowService(ctx, c)
	// init req and assert value
	req := &follow.FollowReq{TargetUser: 2, IsFollow: true}
	resp, err := s.Run(req)

	if err != nil {
		t.Log(err)
	} else {
		assert.NotEqual(t, nil, resp)
	}
}
