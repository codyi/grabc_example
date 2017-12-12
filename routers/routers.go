package routers

import (
	"github.com/astaxie/beego"
	"github.com/grabc"
	"grabc_example/controllers"
	"grabc_example/libs"
	"grabc_example/models"
)

func init() {
	var c []beego.ControllerInterface
	c = append(c, &controllers.SiteController{}, &controllers.UserController{})
	beego.Router("/", &controllers.SiteController{})

	for _, v := range c {
		//将路由注册到beego
		beego.AutoRouter(v)
		//将路由注册到grabc
		grabc.RegisterController(v)
	}

	//注册用户系统模型到grabc
	grabc.RegisterUserModel(&models.User{})

	//忽律权限检查的页面
	grabc.AppendIgnoreRoute("site", "login")

	grabc.Http_403("/site/nopermission")

	grabc.SetLayout(libs.Grabc_layout, nil)
}
