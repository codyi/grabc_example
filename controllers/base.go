package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/codyi/grabc"
	"github.com/codyi/grabc/libs"
	. "grabc_example/models"
	"strings"
	"time"
)

//BaseController
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           User
	libs.Alert
	libs.Breadcrumbs
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

	if this.IsLogin() {
		this.Data["menus"] = libs.ShowMenu(this.controllerName, this.actionName)
		this.Data["alert"] = this.ShowAlert()
		this.Data["breadcrumbs"] = this.ShowBreadcrumbs()
		this.Data["nowTime"] = time.Now().Format("2006-01-02 15:04:05")
		this.Data["user_name"] = this.user.RealName
		this.Data["user_phone"] = this.user.Phone

		//如果grabc采用了新的模板，并且需要在模板中显示数据，可以通过如下的方式添加数据
		grabc.AddLayoutData("user_name", this.Data["user_name"])
		grabc.AddLayoutData("nowTime", this.Data["nowTime"])
		grabc.AddLayoutData("user_phone", this.Data["user_phone"])
	}
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

	grabc.SetBeegoController(&this.Controller)

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
