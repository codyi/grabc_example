package routers

import (
	"github.com/astaxie/beego"
	"goCronJob/controllers"
	"grabc"
)

func init() {
	beego.Router("/", &controllers.SiteController{})
	beego.AutoRouter(&controllers.SiteController{})
	beego.AutoRouter(&controllers.InstallController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.JobController{})

	grabc.RegisterController(&controllers.SiteController{}, &controllers.InstallController{}, &controllers.UserController{}, &controllers.JobController{})
	rbac := grabc.Rbac{}
	rbac.CheckAccess()
}
