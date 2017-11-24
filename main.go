package main

import (
	"github.com/astaxie/beego"
	_ "goCronJob/routers"
)

func main() {
	beego.Run()
}
