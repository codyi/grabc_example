package controllers

type UserController struct {
	BaseController
}

//用户修改密码
func (this *UserController) Modifypassword() {
	this.Layout = "layout/main.html"

	if this.IsPost() {
		this.AddErrorMessage("修改失败")
	}

	this.ShowHtml()
}
