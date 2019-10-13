package modules

import (
	"github.com/doge-soft/dogego_module_mutex"
	"status-monitor/cache"
)

var RedisMutex *dogego_module_mutex.RedisMutex

func InitRedisMutex() {
	RedisMutex = dogego_module_mutex.NewRedisMutex(cache.CacheClient)
}
