package controllers

import (
	"github.com/astaxie/beego"
	"goCronJob/models"
	"strings"
)

type SiteController struct {
	BaseController
}

//Login page
func (this *SiteController) Login() {
	this.setLayout("layout/main-login.html")
	user := models.User{}
	user.ModifyPassword("asdfadsf")
	if this.isPost() {
		phone := strings.TrimSpace(this.GetString("phone"))
		password := strings.TrimSpace(this.GetString("password"))

		user := models.User{}
		err := user.FindByPhone(phone)

		if err != nil {
			this.Data["error"] = err.Error()
		} else if user.ValidatePassword(password) {
			this.redirect(beego.URLFor("SiteController.Index"))
		} else {
			this.Data["error"] = "登录失败"
		}
	}

	this.SetSession("my_session", "asdfads232323f")
	this.showHtml()
}

//main page
func (this *SiteController) Index() {
	this.Data["sess"] = this.GetSession("my_session")
	this.showHtml()
}
