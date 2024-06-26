package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gull28/message_ms/internal/delivery"
	"github.com/gull28/message_ms/internal/models"
)

type Response struct {
	Message string `json:"message"`
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{Message: message}
	json.NewEncoder(w).Encode(response)
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, http.StatusOK, "Messages fetched successfully")
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, http.StatusOK, "Hello, World!")
}

func (app *application) sendCode(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		writeJSONResponse(w, http.StatusBadRequest, "Missing userId")
		return
	}

	codeConfig := app.config.CodeSettings // Load config once

	count, err := models.GetResendCount(app.db, userId)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, "Error getting attempt count")
		return
	}

	if count >= codeConfig.Resends {
		writeJSONResponse(w, http.StatusForbidden, "Try again later!")
		return
	}

	msgType := r.URL.Query().Get("type")
	switch msgType {
	case "phone":
		app.handlePhoneCode(w, r, userId)
	case "email":
		app.handleEmailCode(w, r, userId)
	default:
		writeJSONResponse(w, http.StatusBadRequest, "Invalid message type")
	}
}

func (app *application) handlePhoneCode(w http.ResponseWriter, r *http.Request, userId string) {
	phone := r.URL.Query().Get("phone")
	if !ValidatePhone(phone) {
		writeJSONResponse(w, http.StatusBadRequest, "Invalid phone number")
		return
	}

	code := GenerateCode()
	body := "Your verification code is: " + code

	if err := delivery.SendSMS(phone, body); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, "Error sending code")
		return
	}

	expiry := time.Now().Add(time.Duration(app.config.CodeSettings.Expiry) * time.Minute)
	codeObj, err := models.CreateCode(app.db, userId, code, expiry)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, "Error creating code")
		return
	}

	if err := models.CreateSMS(app.db, userId, phone, codeObj.ID); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, "Error saving SMS record")
		return
	}

	writeJSONResponse(w, http.StatusOK, "Code sent")
}

func (app *application) handleEmailCode(w http.ResponseWriter, r *http.Request, userId string) {
	email := r.URL.Query().Get("email")
	if !ValidateEmail(email) {
		writeJSONResponse(w, http.StatusBadRequest, "Invalid email address")
		return
	}

	code := GenerateCode()
	if err := delivery.SendMail(email, "Verification Code", code); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, "Error sending code")
		return
	}

	expiry := time.Now().Add(time.Duration(app.config.CodeSettings.Expiry) * time.Minute)
	codeObj, err := models.CreateCode(app.db, userId, code, expiry)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, "Error creating code")
		return
	}

	if err := models.CreateMail(app.db, userId, email, codeObj.ID); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, "Error saving mail record")
		return
	}

	writeJSONResponse(w, http.StatusOK, "Code sent")
}

func (app *application) validateCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	userId := r.URL.Query().Get("userId")
	if code == "" || userId == "" {
		writeJSONResponse(w, http.StatusBadRequest, "Missing code or userId")
		return
	}

	isValid, err := models.CheckValidity(app.db, code, userId)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, "Error validating code")
		return
	}

	if !isValid {
		writeJSONResponse(w, http.StatusForbidden, "Invalid code")
		return
	}

	writeJSONResponse(w, http.StatusOK, "Code is valid")
}
