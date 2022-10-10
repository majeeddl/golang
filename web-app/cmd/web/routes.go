package main

import (
	"github.com/go-chi/chi"
	"webapp/pkg/handlers"

	//"github.com/bmizerany/pat"
	"net/http"
	//"webapp/pkg/config"
)

func routes() http.Handler {
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.About))

	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	return mux
}
