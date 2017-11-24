package controllers

import ()

type SiteController struct {
	BaseController
}

//Login page
func (this *SiteController) Login() {
	this.setLayout("layout/main-login.html")
	this.SetSession("my_session", "asdfads232323f")
	this.showHtml()
}

//main page
func (this *SiteController) Index() {
	this.Data["sess"] = this.GetSession("my_session")
	this.showHtml()
}
