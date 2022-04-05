package main

import (
	_ "coinwallet/common/template"
	"coinwallet/cron"
	_ "coinwallet/routers"
	_ "coinwallet/session"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	cron.Run()
	beego.Run()
}

