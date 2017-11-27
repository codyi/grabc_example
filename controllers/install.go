package controllers

type InstallController struct {
	BaseController
}

//User install page
func (this *InstallController) Index() {
	this.showHtml()
}
