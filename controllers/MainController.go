package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Websitr"] = "My first web"

	this.show()
}
