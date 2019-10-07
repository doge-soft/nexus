package conf

import (
	"doge-service-status/inits"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	// 初始化
	inits.Init()
}
