package main

import (
	"github.com/astaxie/beego"
	"github.com/codyi/grabc"
	"grabc_example/controllers"
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

	//增加忽律权限检查的页面
	grabc.AppendIgnoreRoute("site", "login")

	//设置权限不足的提示页面路径
	grabc.Http_403("/site/nopermission")

	//设置grabc页面路径
	//如果使用默认的，不要设置或者置空
	// grabc.SetViewPath("")

	//设置grabc的layout
	// menus := make(map[string]interface{}, 0)
	// menus["grabc_menus"] = grabc.AccessMenus()
	// grabc.SetLayout("layout/main.html", menus)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "/tmp"
	beego.Run()
}
