package controllers

type UserController struct {
	BaseController
}

//用户修改密码
func (this *UserController) Modifypassword() {
	this.setPageTitle("修改密码")
	this.addBreadcrumbs("管理员", this.URLFor("UserController.Index"))
	this.addBreadcrumbs("修改密码", "")

	this.showHtml()
}
