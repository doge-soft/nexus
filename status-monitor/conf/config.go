package conf

import (
	"status-monitor/inits"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	// 初始化
	inits.Init()
}
