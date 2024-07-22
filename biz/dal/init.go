package dal

import (
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
