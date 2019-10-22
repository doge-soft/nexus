package main

import (
	_ "nexus/auth-gateway/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
