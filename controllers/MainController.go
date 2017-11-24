package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Websitr"] = "My first web"
	this.Data["Email"] = "mail@liguosong.com"
	this.TplName = "index.tpl"
}
