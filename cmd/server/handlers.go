package main

import (
	"net/http"
)

func getMessages(r *http.Request, w http.ResponseWriter) {
	validateMethod(r, w, http.MethodGet)

	w.Header().Set("Content-Type", "application/json")
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if validateMethod(r, w, http.MethodGet) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello, World!"}`))

		return
	}
}
