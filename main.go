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
	beego.Router("/", &controllers.SiteController{}, "*:Index")

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

	//设置grabc页面路径
	//如果使用默认的，不要设置或者置空
	//注意：设置grabc的模板必须在beego.Run()之前设置，如果视图目录在当前项目中，可以使用相对目录，否则需要绝对路径
	// grabc.SetViewPath("views")

	//设置grabc的layout
	grabc.SetLayout("layout/main.html", "views")

	//注册获取当前登录用户ID的函数
	grabc.RegisterUserIdFunc(func(c *beego.Controller) int {
		sessionUId := c.GetSession("login_user_id")

		if sessionUId != nil {
			user := models.User{}
			user.FindById(sessionUId.(int))
			return user.Id
		}

		return 0
	})
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "/tmp"
	beego.Run()
}
