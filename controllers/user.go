package controllers

type UserController struct {
	BaseController
}

//用户修改密码
func (this *UserController) Modifypassword() {
	this.Layout = "layout/main.html"
	this.ShowHtml()
}
