package routers

import (
	"github.com/astaxie/beego"
	"goCronJob/controllers"
)

func init() {
	beego.AutoRouter(&controllers.SiteController{})
	beego.AutoRouter(&controllers.InstallController{})
	beego.AutoRouter(&controllers.UserController{})
}
