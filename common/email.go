package common

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"strings"
)

// SendEmail 发送验证邮件
func SendEmail(code string, fem, tem, emCode string) {
	e := email.NewEmail()
	e.From = fem
	e.To = []string{tem}
	// 设置主题
	e.Subject = "邮箱验证"
	// 设置主体
	e.Text = []byte(code)
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", fem, emCode, "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}

// GenValidateCode 生成验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
