package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type SiteController struct {
	BaseController
}

//Login page
func (this *SiteController) Login() {
	this.setLayout("layout/main-login.html")

	if this.isLogin() {
		this.redirect(beego.URLFor("SiteController.Index"))
	}

	if this.isPost() {
		phone := strings.TrimSpace(this.GetString("phone"))
		password := strings.TrimSpace(this.GetString("password"))

		if err := this.BaseController.login(phone, password); err == nil {
			this.redirect(beego.URLFor("SiteController.Index"))
		} else {
			this.Data["error"] = err.Error()
		}
	}

	this.showHtml()
}

//main page
func (this *SiteController) Index() {
	this.showHtml()
}

func (this *SiteController) Logout() {
	fmt.Println("loginout")
	if this.isLogin() {
		this.BaseController.logout()
	}

	this.redirect(beego.URLFor("SiteController.Login"))
}
