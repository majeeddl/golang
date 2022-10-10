package main

import (
	"fmt"
	"net/http"
)

func WriteToConsoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" HIT THE PAGE")
		next.ServeHTTP(w, r)
	})
}
