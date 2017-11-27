package libs

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"strconv"
)

//check is application is installed
func IsInstall() bool {
	isInstall, _ := strconv.ParseBool(beego.AppConfig.String("isInstall"))
	return isInstall
}

//md5 input string
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
