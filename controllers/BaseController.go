package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

func init() {

}

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (self *BaseController) Prepare() {
	controlerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.setLayout("layout/main.html")
}

// redirect to url
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

// load tpl
func (self *BaseController) show(tpl ...string) {
	if len(tpl) > 0 {
		self.TplName = tpl[0]
	} else {
		self.TplName = self.controllerName + "/" + self.actionName + ".html"
	}
}

//set layout
func (self *BaseController) setLayout(layout string) {
	self.Layout = layout
}
