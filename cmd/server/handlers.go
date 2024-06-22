package main

import (
	"net/http"

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

	code := GenerateCode()

	msgType := r.URL.Query().Get("type")

	if msgType == "phone" {
		if ValidatePhone(r.URL.Query().Get("phone")) == false {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Invalid phone number!"}`))
			return
		}

		count, err := models.GetSmsAttemptCount(app.db, r.URL.Query().Get("userId"))

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Error getting attempt count!"}`))
			return
		}

		// todo: get config value from .env

	}

	if msgType == "email" {
		// validate email
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Code sent!"}`))

	return
}

func (app *application) validateCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Code validated!"}`))

	return
}
