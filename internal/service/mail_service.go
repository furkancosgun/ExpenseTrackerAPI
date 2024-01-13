package service

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/furkancosgun/expense-tracker-api/config/mail"
)

type MailContent struct {
	To      string
	Subject string
	Body    string
	Config  mail.Config
}

func buildMail(mail MailContent) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Config.Username)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join([]string{mail.To}, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func NewOtpMailContent(to string, otp string, expiresAt time.Time) *MailContent {
	var newConfig = mail.NewCofig()

	//Subject
	var subject = fmt.Sprintf("%s: %s", "Expense-Tracker Account Verification OTP Code", otp)
	var content = `
	<html> <body> <div style=" font-family: Helvetica, Arial, sans-serif; min-width: 1000px; overflow: auto; line-height: 2; " > <div style="margin: 50px auto; width: 70%; padding: 20px 0"> <div style="border-bottom: 1px solid #eee"> <a href="" style=" font-size: 1.4em; color: #00466a; text-decoration: none; font-weight: 600; " >Expense Tracker</a > </div> <p style="font-size: 1.1em">Hi,</p> <p>Thank you for choosing Expense Tracker. Use the following OTP Code.</p> <b>Expires At: {{EXPIRES_AT}}</b> <h2 style=" background: #00466a; margin: 0 auto; width: max-content; padding: 0 10px; color: #fff; border-radius: 4px; " > {{OTP_CODE}} </h2> <p style="font-size: 0.9em">Regards,<br />Expense Tracker</p> <hr style="border: none; border-top: 1px solid #eee" /> <div style=" float: right; padding: 8px 0; color: #aaa; font-size: 0.8em; line-height: 1; font-weight: 300; " ></div> </div> </div> </body></html>
`

	var body = strings.Replace(content, "{{OTP_CODE}}", otp, 1)
	body = strings.Replace(body, "{{EXPIRES_AT}}", expiresAt.Format("2006-01-02 15:04:05"), 1)

	return &MailContent{To: to, Subject: subject, Body: body, Config: *newConfig}
}

func (mail *MailContent) Send() error {

	addr := fmt.Sprintf("%s:%s", mail.Config.Host, mail.Config.Port)
	auth := smtp.PlainAuth(mail.Config.Identity, mail.Config.Username, mail.Config.Password, mail.Config.Host)
	msgString := buildMail(*mail)
	return smtp.SendMail(
		addr,
		auth,
		mail.Config.Username,
		[]string{mail.To},
		[]byte(msgString),
	)
}
