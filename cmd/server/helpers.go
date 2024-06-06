package main

import (
	"net/http"
)

func validateMethod(r *http.Request, w http.ResponseWriter, method string) bool {
	if r.Method != method {
		w.Header().Set("Allow", method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return false
	}

	return true
}

func validateBearerToken(r *http.Request, w http.ResponseWriter) bool {
	if r.Header.Get("Authorization") == "" {

		http.Error(w, "Unauthorized", http.StatusUnauthorized)

		return false
	}

	// todo: implement bearer token validation

	return true
}
