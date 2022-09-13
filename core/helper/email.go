package helper

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

// 邮箱验证码发送
func MailSendCode(mail,code string) error{
	e := email.NewEmail()
	// 1. 发送者姓名、邮件
	e.From = "若水 <1085252985@qq.com>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为<h1>"+code+"</h1>")

	err := e.SendWithTLS("smtp.qq.com:25", smtp.PlainAuth("", "1085252985@qq.com", "ccybguhvospobaad", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify:true,ServerName:"smtp.qq.com"})
	if err!=nil{
		return err
	}
	return nil
}


func RandCode() string{
	// 随机6次一位数，组合成一个6位数
	s := "1234567890"
	rand.Seed(time.Now().UnixNano())

	code := ""
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}

	return code
}