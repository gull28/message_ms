package main

import (
	"net/http"
	"time"

	"github.com/gull28/message_ms/internal/config"
	"github.com/gull28/message_ms/internal/delivery"
	"github.com/gull28/message_ms/internal/models"
)

func getMessages(r *http.Request, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello, World!"}`))

	return
}

func (app *application) sendCode(w http.ResponseWriter, r *http.Request) {
	count, err := models.GetResendCount(app.db, r.URL.Query().Get("userId"))

	codeConfig := config.LoadConfig().CodeSettings

	if count >= codeConfig.Resends {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message": "Try again later!"}`))
		return
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Error getting attempt count!"}`))
		return
	}

	msgType := r.URL.Query().Get("type")

	if msgType == "phone" {
		if ValidatePhone(r.URL.Query().Get("phone")) == false {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Invalid phone number!"}`))
			return
		}

		code := GenerateCode()

		body := "Your verification code is: " + code
		err := delivery.SendSMS(r.URL.Query().Get("phone"), body)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Error sending code!"}`))
			return
		}

		codeObj, err := models.CreateCode(app.db, r.URL.Query().Get("userId"), code, time.Now().Add(time.Duration(codeConfig.Expiry)*time.Minute))

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Error creating code!"}`))
			return
		}

		err = models.CreateSMS(app.db, r.URL.Query().Get("userId"), r.URL.Query().Get("phone"), codeObj.ID)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Error sending code!"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Code sent!"}`))
		return
	}

	if msgType == "email" {
		if ValidateEmail(r.URL.Query().Get("email")) == false {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Invalid email address!"}`))
			return
		}

		code := GenerateCode()

		err := delivery.SendMail(r.URL.Query().Get("email"), "Verification Code", code)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Error sending code!"}`))
			return
		}

		codeObj, err := models.CreateCode(app.db, r.URL.Query().Get("userId"), code, time.Now().Add(time.Duration(codeConfig.Expiry)*time.Minute))

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Error creating code!"}`))
			return
		}

		err = models.CreateMail(app.db, r.URL.Query().Get("userId"), r.URL.Query().Get("email"), codeObj.ID)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Error creating mail!"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Code sent!"}`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Code sent!"}`))

	return
}

func (app *application) validateCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	userId := r.URL.Query().Get("userId")

	isValid, err := models.CheckValidity(app.db, code, userId)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Error validating code!"}`))
		return
	}

	if isValid == false {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Invalid code!"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Code is valid!"}`))
}
