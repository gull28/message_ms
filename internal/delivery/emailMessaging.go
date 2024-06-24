package delivery

import (
	"net/smtp"

	"github.com/gull28/message_ms/internal/config"
)

func sendEmail(to string, subject string, body string) error {
	smtpConfig := config.LoadConfig().Smtp

	smtpHost := smtpConfig.Host
	smtpPort := smtpConfig.Port
	smtpUser := smtpConfig.Username
	smtpPassword := smtpConfig.Password

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

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
