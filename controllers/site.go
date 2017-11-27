package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type SiteController struct {
	BaseController
}

//Login page
func (this *SiteController) Login() {
	this.setLayout("layout/main-login.html")

	if this.isPost() {
		phone := strings.TrimSpace(this.GetString("phone"))
		password := strings.TrimSpace(this.GetString("password"))

		if err := this.BaseController.Login(phone, password); err == nil {
			this.redirect(beego.URLFor("SiteController.Index"))
		} else {
			this.Data["error"] = err.Error()
		}
	}

	this.showHtml()
}

//main page
func (this *SiteController) Index() {
	this.Data["sess"] = this.GetSession("")
	this.showHtml()
}
