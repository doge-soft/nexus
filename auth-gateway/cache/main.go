package cache

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func ConnectRedisCache() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis_addr"),
		Password: beego.AppConfig.String("redis_password"),
		DB:       beego.AppConfig.DefaultInt("redis_db", 0),
	})
}
