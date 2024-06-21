package main

import (
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Code sent!"}`))

	return
}

func (app *application) validateCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Code validated!"}`))

	return
}
