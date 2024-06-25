package delivery

import (
	"net/smtp"

	"github.com/gull28/message_ms/internal/config"
)

func SendMail(to string, subject string, code string) error {
	smtpConfig := config.LoadConfig().Smtp

	smtpHost := smtpConfig.Host
	smtpPort := smtpConfig.Port
	smtpUser := smtpConfig.Username
	smtpPassword := smtpConfig.Password

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

	body := BuildEmailMessage(code)
	// body := "Your verification code is: " + code
	from := smtpUser
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	recipient := []string{to}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipient, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func BuildEmailMessage(code string) string {
	// implement your own email message for production
	// use a template engine or a library to build the message for convenience
	return "Your verification code is: " + code
}
