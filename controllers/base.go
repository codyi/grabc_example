package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"goCronJob/libs"
	. "goCronJob/models"
	"strings"
)

//BaseController
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	siteTitle      string
	user           User
}

// Prepare
func (this *BaseController) Prepare() {
	controlerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.setLayout("layout/main.html")
	this.siteTitle = beego.AppConfig.String("site.title")
	this.checkPermision()
}

// redirect to url
func (this *BaseController) redirect(url string) {
	// this.Redirect(url, 301)
	this.Controller.Redirect(url, 301)
	this.StopRun()
}

// load tpl
func (this *BaseController) showHtml(tpl ...string) {
	if len(tpl) > 0 {
		this.TplName = tpl[0]
	} else {
		this.TplName = this.controllerName + "/" + this.actionName + ".html"
	}

	this.Data["siteTitle"] = this.siteTitle
	this.Data["siteStaticVersion"] = beego.AppConfig.String("site.static.version")
	this.Data["app_name"] = beego.AppConfig.String("appname")
}

//set layout
func (this *BaseController) setLayout(layout string) {
	this.Layout = layout
}

//set web title
func (this *BaseController) setSiteTile(title string) {
	this.siteTitle = title
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//login by phone and password
func (this *BaseController) login(phone string, password string) (err error) {
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
func (this *BaseController) logout() {
	this.DelSession("login_user_id")
}

//check is user login
func (this *BaseController) isLogin() bool {
	return this.user.Id > 0
}

//check user permission
func (this *BaseController) checkPermision() {
	if !libs.IsInstall() {
		this.checkInstall()
		return
	}

	sessionUId := this.GetSession("login_user_id")

	if sessionUId != nil {
		this.user = User{}
		this.user.FindById(sessionUId.(int))
	}

	if this.isLogin() {
		//todo: check user pemission
	} else if (this.controllerName != "site" && this.actionName != "login") || (this.controllerName == "site" && this.actionName != "login") {
		this.redirect(beego.URLFor("SiteController.Login"))
	}
}

////if application is not installed,will redirect to install page
func (this *BaseController) checkInstall() {
	if !libs.IsInstall() && this.controllerName != "install" {
		this.redirect(beego.URLFor("InstallController.Index"))
	}
}
