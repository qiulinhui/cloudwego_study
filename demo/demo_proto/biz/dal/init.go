package dal

import (
	"cloudwego_study/demo/demo_proto/biz/dal/mysql"
	"cloudwego_study/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
