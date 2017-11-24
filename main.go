package main

import (
	"github.com/astaxie/beego"
	_ "goCronJob/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "/tmp"

	beego.Run()
}
