package main

import (
	"github.com/astaxie/beego"
	_ "grabc_example/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "/tmp"
	beego.Run()
}
