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
	if this.isLogin() {
		this.BaseController.logout()
	}

	this.redirect(beego.URLFor("SiteController.Login"))
}

//main page
func (this *SiteController) NoPermission() {
	this.setPageTitle("权限不够")
	this.addBreadcrumbs("错误提示", "")
	this.showHtml("site/403.html")
}

func (this *SiteController) RABCMethods() []string {
	return []string{"Get", "Post"}
}
