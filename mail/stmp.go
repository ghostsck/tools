package mail

import (
	"tools/setting"
	"fmt"
	"gopkg.in/gomail.v2"
)

type Email struct {
	to      string "to"
	subject string "subject"
	body    string "msg"
}

func NewEmail(to, subject, body string) *Email {
	return &Email{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (email *Email) Send() (bool, error) {

	m := gomail.NewMessage()

	// 发件人
	m.SetAddressHeader("From", setting.MAIL_FROM_ADDRESS, setting.MAIL_FROM_NAME)

	// 收件人
	m.SetHeader("To", m.FormatAddress(email.to, email.to))
	//抄送
	//m.SetHeader("Cc", m.FormatAddress("xxxx@foxmail.com", "收件人"))
	//暗送
	//m.SetHeader("Bcc", m.FormatAddress("xxxx@gmail.com", "收件人"))
	// 主题
	m.SetHeader("Subject", email.subject)

	// 可以放html..还有其他的
	m.SetBody("text/html", email.body)

	//添加附件
	//m.Attach("我是附件")

	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewDialer(setting.MAIL_HOST, setting.MAIL_PORT, setting.MAIL_USERNAME, setting.MAIL_PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("发送失败", err)
		return false, err
	}

	return true, nil
}
