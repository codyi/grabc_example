package controllers

import (
	"strings"
)

type UserController struct {
	BaseController
}

//用户修改密码
func (this *UserController) Modifypassword() {
	this.setPageTitle("修改密码")
	this.addBreadcrumbs("管理员", this.URLFor("UserController.Index"))
	this.addBreadcrumbs("修改密码", "")

	if this.isPost() {
		old_password := strings.TrimSpace(this.GetString("old_password"))
		new_password := strings.TrimSpace(this.GetString("new_password"))
		confirm_password := strings.TrimSpace(this.GetString("confirm_password"))

		if this.user.ValidatePassword(old_password) {
			if new_password == confirm_password {
				if this.user.ModifyPassword(new_password) {
					this.AddSuccessMessage("修改成功")
				} else {
					this.AddErrorMessage("修改失败")
				}
			} else {
				this.AddErrorMessage("新密码与确认密码不匹配")
			}
		} else {
			this.AddErrorMessage("原始密码不正确")
		}
	}

	this.showHtml()
}

//管理员列表
func (this *UserController) Index() {
	this.setPageTitle("管理员列表")
	this.addBreadcrumbs("管理员列表", "")

	this.showHtml()
}
