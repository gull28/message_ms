package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
)

// func validateMethod(r *http.Request, w http.ResponseWriter, method string) bool {
// 	if r.Method != method {
// 		w.Header().Set("Allow", method)
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

// 		return false
// 	}

// 	return true
// }

// func validateBearerToken(r *http.Request, w http.ResponseWriter) bool {
// 	if r.Header.Get("Authorization") == "" ||  {

// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)

// 		return false
// 	}

// 	// todo: implement bearer token validation

// 	return true
// }

func MultipleMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func GenerateCode() string {
	code := rand.Intn(999999)

	return fmt.Sprintf("%06d", code)
}

// validate phone number format (e.g. +1234567890)
// feel free to customize the pattern to your needs as this is just a simple regex validation
func ValidatePhone(phone string) bool {
	pattern := `^\+\d{1,3}\d{1,4}\d{1,4}\d{1,4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(phone)
}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}
