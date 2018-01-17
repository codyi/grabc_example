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
	this.Layout = "layout/login.html"

	if this.IsLogin() {
		this.Redirect(beego.URLFor("SiteController.Index"))
	}

	if this.IsPost() {
		phone := strings.TrimSpace(this.GetString("phone"))
		password := strings.TrimSpace(this.GetString("password"))

		if err := this.BaseController.Login(phone, password); err == nil {
			this.Redirect(beego.URLFor("SiteController.Index"))
		} else {
			this.Data["error"] = err.Error()
		}
	}

	this.ShowHtml()
}

//main page
func (this *SiteController) Index() {
	this.Layout = "layout/main.html"
	this.ShowHtml()
}

func (this *SiteController) Logout() {
	if this.IsLogin() {
		this.BaseController.Logout()
	}

	this.Redirect(beego.URLFor("SiteController.Login"))
}

//main page
func (this *SiteController) NoPermission() {
	this.Layout = "layout/main.html"
	this.ShowHtml("site/403.html")
}

func (this *SiteController) RABCMethods() []string {
	return []string{"Get", "Post"}
}
