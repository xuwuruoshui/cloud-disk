package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

// 要多发送几次，可能发送失败
func TestEamil(t *testing.T){
	e := email.NewEmail()
	// 1. 发送者姓名、邮件
	e.From = "若水 <1085252985@qq.com>"
	e.To = []string{"xuwuruoshui@gmail.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为<h1>666666</h1>")

	err := e.SendWithTLS("smtp.qq.com:25", smtp.PlainAuth("", "1085252985@qq.com", "ccybguhvospobaad", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify:true,ServerName:"smtp.qq.com"})
	if err!=nil{
		t.Fatal(err)
	}
}
