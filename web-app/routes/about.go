package routes

import (
	"net/http"
)

func AboutRoutes(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "About Page Web App")
	renderTemplate(w, "about.page.tmpl")
}
