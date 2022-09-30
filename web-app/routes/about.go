package routes

import (
	"net/http"
	"webapp/pkg/render"
)

func AboutRoutes(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "About Page Web App")
	render.RenderTemplate(w, "about.page.tmpl")
}
