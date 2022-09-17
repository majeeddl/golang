package routes

import (
	"fmt"
	"net/http"
)

func AboutRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page Web App")
}
