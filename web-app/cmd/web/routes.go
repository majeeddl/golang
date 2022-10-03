package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"webapp/pkg/handlers"
	//"webapp/pkg/config"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/about", http.HandlerFunc(handlers.About))

	return mux
}
