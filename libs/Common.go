package libs

import (
	"github.com/astaxie/beego"
	"strconv"
)

//check is application is installed
func IsInstall() bool {
	isInstall, _ := strconv.ParseBool(beego.AppConfig.String("isInstall"))
	return isInstall
}
