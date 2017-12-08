package libs

import (
	"crypto/md5"
	"encoding/hex"
)

//md5 input string
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
