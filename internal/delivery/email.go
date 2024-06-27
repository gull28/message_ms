package delivery

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/gull28/message_ms/internal/config"
)

// SendMail sends an email with an HTML body loaded from a template file
func SendMail(to string, subject string, code string) error {
	// Load SMTP configuration
	smtpConfig := config.LoadConfig().Smtp
	smtpHost := smtpConfig.Host
	smtpPort := smtpConfig.Port
	smtpUser := smtpConfig.Username
	smtpPassword := smtpConfig.Password

	htmlBody, err := LoadHTMLTemplate("./templates/index.html", code)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

	from := smtpUser
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-Version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s",
		from, to, subject, htmlBody)

	recipient := []string{to}
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipient, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func LoadHTMLTemplate(filePath string, code string) (string, error) {
	htmlData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	htmlContent := string(htmlData)

	htmlContent = strings.Replace(htmlContent, "{{.Code}}", code, -1)

	return htmlContent, nil
}
