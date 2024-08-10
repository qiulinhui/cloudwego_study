package dal

import (
	"cloudwego_study/demo/demo_thrift/biz/dal/mysql"
	"cloudwego_study/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
