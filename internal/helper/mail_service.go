package helper

import (
	"fmt"
	"net/smtp"

	"github.com/furkancosgun/expense-tracker-api/config/mail"
)

type MailContent struct {
	To      string
	Subject string
	Body    string
	Config  mail.Config
}

func NewMailContent(to string, subject string, body string) *MailContent {
	var newConfig = mail.NewCofig()
	return &MailContent{To: to, Subject: subject, Body: body, Config: *newConfig}
}

func (mail *MailContent) Send() error {

	addr := fmt.Sprintf("%s:%s", mail.Config.Host, mail.Config.Port)
	auth := smtp.PlainAuth(mail.Config.Identity, mail.Config.Username, mail.Config.Password, mail.Config.Host)

	return smtp.SendMail(
		addr,
		auth,
		mail.Config.Username,
		[]string{mail.To},
		[]byte(mail.Body),
	)
}
