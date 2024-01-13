package service

import (
	"fmt"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"

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

func getCurrentDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	currentFileDir := filepath.Dir(ex)
	return currentFileDir, nil
}

func NewOtpMailContent(to string, otp string) *MailContent {
	var newConfig = mail.NewCofig()

	//Subject
	var subject = fmt.Sprintf("%s: %s", "OTP CODE", otp)

	//Content
	currentDir, _ := getCurrentDir()
	var contentPath = fmt.Sprintf("%s/%s", currentDir, "../internal/resource/otp_mail_content.html")
	content, _ := os.ReadFile(contentPath)
	body := string(content)
	body = strings.Replace(body, "{{OTP_CODE}}", otp, 1)

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
