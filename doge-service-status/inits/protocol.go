package inits

import "doge-service-status/web/servers"

// 初始化服务器级别的协议
func InitProtocol() {
	servers.RegisterServers()
}
