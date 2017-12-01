package controllers

type JobController struct {
	BaseController
}

func (this *JobController) Index() {
	this.showHtml()
}
