// Package public 公共目录
package public

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

// Md5Encoder MD5加密
func Md5Encoder(code string) string {
	m := md5.New()
	_, _ = io.WriteString(m, code)
	return hex.EncodeToString(m.Sum(nil))

}

// SaltPassword 密码加盐
func SaltPassword(pwd, salt string) string {
	saltPwd := fmt.Sprintf("%s$%s", Md5Encoder(pwd), salt)
	return saltPwd
}
