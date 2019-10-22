package main

import (
	_ "nexus/auth-gateway/routers"
	_ "nexus/auth-gateway/conf"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
