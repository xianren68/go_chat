package common

import (
	"fmt"
	"testing"
)

func TestSendEmail(t *testing.T) {
	SendEmail("1234", "xianren711@qq.com", "xianren711@gmail.com", "fddlauhpzdkcdehg")
}

func TestGenValidateCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GenValidateCode(6))
	}
}
