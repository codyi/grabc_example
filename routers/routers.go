package routers

import (
	"github.com/astaxie/beego"
	"goCronJob/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
