package controllers

type InstallController struct {
	BaseController
}

//User install page
func (this *InstallController) Index() {
	this.setLayout("layout/main-login.html")
	this.showHtml()
}
