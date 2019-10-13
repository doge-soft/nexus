package modules

import (
	"github.com/doge-soft/dogego_module_mq"
	"status-monitor/cache"
)

var RedisMQ *dogego_module_mq.RedisMQ

func InitRedisMQModule() {
	RedisMQ = dogego_module_mq.NewRedisMQ(cache.CacheClient)
}
