package dal

import (
	"github.com/qiulinhui/cloudwego_study/biz/dal/mysql"
	"github.com/qiulinhui/cloudwego_study/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
