package delivery

import (
	"github.com/gull28/message_ms/internal/config"
	// "github.com/sfreiberg/gotwilio"
)

func SendSMS(to string, body string) error {
	smsConfig := config.LoadConfig().Sms

	host := smsConfig.SMSHost
	apiKey := smsConfig.APIKey
	senderID := smsConfig.SenderID
	from := smsConfig.From

	// implement your own sms sending logic suited to your api, return an error if failed, simple as that
	// if you use twilio, simply uncomment the import and use the api from the library

	return nil
}
