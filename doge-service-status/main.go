package main

import (
	_ "doge-service-status/conf"
	"doge-service-status/web/routers"
	"doge-service-status/web/servers"
	"sync"
)

func main() {
	router := routers.NewRouter()
	group := sync.WaitGroup{}

	// 启动所有注册过的服务器
	servers.StartServers(&group, router)
}
