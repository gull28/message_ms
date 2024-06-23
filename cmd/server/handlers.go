package main

import (
	"net/http"

	"github.com/gull28/message_ms/internal/config"
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
	count, err := models.GetAttemptCount(app.db, r.URL.Query().Get("userId"))

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

	}

	if msgType == "email" {
		if ValidateEmail(r.URL.Query().Get("email")) == false {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Invalid email address!"}`))
			return
		}

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
