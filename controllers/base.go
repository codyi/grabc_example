package controllers

import (
	"github.com/astaxie/beego"
	"goCronJob/libs"
	"strings"
)

//BaseController
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	siteTitle      string
}

// Prepare
func (self *BaseController) Prepare() {
	controlerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.setLayout("layout/main.html")
	self.siteTitle = beego.AppConfig.String("site.title")

	//if application is not installed,will redirect to install page
	if libs.IsInstall() {

	} else if self.controllerName != "install" {
		self.controllerName = "install"
		self.actionName = "index"
	}
}

// redirect to url
func (self *BaseController) redirect(url string) {
	// self.Redirect(url, 301)
	self.Controller.Redirect(url, 301)
	self.StopRun()
}

// load tpl
func (self *BaseController) showHtml(tpl ...string) {
	if len(tpl) > 0 {
		self.TplName = tpl[0]
	} else {
		self.TplName = self.controllerName + "/" + self.actionName + ".html"
	}

	self.Data["siteTitle"] = self.siteTitle
	self.Data["siteStaticVersion"] = beego.AppConfig.String("site.static.version")
	self.Data["app_name"] = beego.AppConfig.String("appname")
}

//set layout
func (self *BaseController) setLayout(layout string) {
	self.Layout = layout
}

//set web title
func (self *BaseController) setSiteTile(title string) {
	self.siteTitle = title
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}
