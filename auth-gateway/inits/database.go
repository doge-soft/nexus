package inits

import (
	"github.com/astaxie/beego"
	"nexus/auth-gateway/cache"
	"nexus/auth-gateway/datasource"
)

func InitDatabase() {
	// 链接mysql数据库
	datasource.ConnectDatabase(
		beego.AppConfig.String("master_dsn"),
		beego.AppConfig.String("slave_dsn"),
	)

	// 连接Redis缓存
	cache.ConnectRedisCache()
}
