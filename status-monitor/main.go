package main

import (
	_ "status-monitor/conf"
	"status-monitor/web/routers"
	"status-monitor/web/servers"
	"sync"
)

func main() {
	router := routers.NewRouter()
	group := sync.WaitGroup{}

	// 启动所有注册过的服务器
	servers.StartServers(&group, router)
}
