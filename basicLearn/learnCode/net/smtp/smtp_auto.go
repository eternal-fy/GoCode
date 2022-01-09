package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

const (
	to                 = "906542557@qq.com"
	SMTP_MAIL_NICKNAME = "eternal-fy"
	SMTP_MAIL_USER     = "liude"
)

func main() {
	body := getBody()
	// Set up authentication information.
	subject := "主题"
	contentType := "Content-Type: text/html; charset=UTF-8"

	auth := smtp.PlainAuth(
		"",
		"906542557@qq.com",
		"zuxncrpzuipwbedd",
		"smtp.qq.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", to, SMTP_MAIL_NICKNAME, SMTP_MAIL_USER, subject, contentType, body)
	msg := []byte(s)
	err := smtp.SendMail(
		"smtp.qq.com:587",
		auth,
		"906542557@qq.com",
		[]string{"906542557@qq.com"},
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}
}
func getBody() string {
	inputFile, _ := os.Open("d:/s.html")

	buf := make([]byte, 1024*100)
	n, _ := inputFile.Read(buf)
	fmt.Println(string(buf[:n]))
	return string(buf[:n])

}

/*
package main

import (
"fmt"
"net/smtp"
)

const (
	// 邮件服务器地址
	SMTP_MAIL_HOST = "smtp.qq.com"
	// 端口
	SMTP_MAIL_PORT = "587"
	// 发送邮件用户账号
	SMTP_MAIL_USER = "1348581672@qq.com"
	// 授权密码
	SMTP_MAIL_PWD = "xxxx"
	// 发送邮件昵称
	SMTP_MAIL_NICKNAME = "lewis"
)

func main() {
	address := []string{"norton_s@qq.com", "lewissunp@outlook.com"}
	subject := "test mail"
	body :=
		`<br>hello!</br>
    <br>this is a test email, pls ignore it.</br>`
	SendMail(address, subject, body)

}

func SendMail(address []string, subject, body string) (err error) {
	// 认证, content-type设置
	auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	contentType := "Content-Type: text/html; charset=UTF-8"

	// 因为收件人，即address有多个，所以需要遍历,挨个发送
	for _, to := range address {
		s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", to, SMTP_MAIL_NICKNAME, SMTP_MAIL_USER, subject, contentType, body)
		msg := []byte(s)
		addr := fmt.Sprintf("%s:%s", SMTP_MAIL_HOST, SMTP_MAIL_PORT)
		err = smtp.SendMail(addr, auth, SMTP_MAIL_USER, []string{to}, msg)
		if err != nil {
			return err
		}
	}
	return err
}*/
