package main

import (
	"net/http"
)

func getMessages(r *http.Request, w http.ResponseWriter) {
	validateMethod(r, w, http.MethodGet)

	w.Header().Set("Content-Type", "application/json")
}
