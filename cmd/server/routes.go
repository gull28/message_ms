package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", MultipleMiddleware(http.HandlerFunc(app.home), AuthMiddleware, MethodMiddleware(http.MethodGet)))

	mux.Handle("/send-code", MultipleMiddleware(http.HandlerFunc(app.sendCode), AuthMiddleware, MethodMiddleware(http.MethodPost)))

	mux.Handle("/verify-code", MultipleMiddleware(http.HandlerFunc(app.validateCode), AuthMiddleware, MethodMiddleware(http.MethodPost)))

	return mux
}
