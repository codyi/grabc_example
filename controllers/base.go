package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/codyi/grabc"
	. "grabc_example/models"
	"strings"
)

//BaseController
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           User
}

// Prepare
func (this *BaseController) Prepare() {
	controlerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.CheckPermision()
}

// redirect to url
func (this *BaseController) Redirect(url string) {
	this.Controller.Redirect(url, 302)
	this.StopRun()
}

// load tpl
func (this *BaseController) ShowHtml(tpl ...string) {
	if len(tpl) > 0 {
		this.TplName = tpl[0]
	} else {
		this.TplName = this.controllerName + "/" + this.actionName + ".html"
	}

	this.Data["grabc_menus"] = grabc.AccessMenus()
}

//login by phone and password
func (this *BaseController) Login(phone string, password string) (err error) {
	if phone == "" || password == "" {
		return errors.New("登录信息不能为空")
	}

	user := User{}
	err = user.FindByPhone(phone)

	if err != nil {
		return err
	} else if user.ValidatePassword(password) {
		this.SetSession("login_user_id", user.Id)
		return nil
	} else {
		return errors.New("登录信息不正确")
	}
}

//user logout
func (this *BaseController) Logout() {
	this.DelSession("login_user_id")
}

//check is user login
func (this *BaseController) IsLogin() bool {
	return this.user.Id > 0
}

//check user permission
func (this *BaseController) CheckPermision() {
	sessionUId := this.GetSession("login_user_id")

	if sessionUId != nil {
		this.user = User{}
		this.user.FindById(sessionUId.(int))
	}

	grabc.RegisterIdentify(this.user)

	if this.IsLogin() {
		if !grabc.CheckAccess(this.controllerName, this.actionName) {
			this.Redirect(this.URLFor("SiteController.NoPermission"))
		}
	} else if (this.controllerName != "site" && this.actionName != "login") || (this.controllerName == "site" && this.actionName != "login") {
		this.Redirect(this.URLFor("SiteController.Login"))
	}
}

// 是否POST提交
func (this *BaseController) IsPost() bool {
	return this.Ctx.Request.Method == "POST"
}
