package routes

import (
	"net/http"
	"webapp/render"
)

func HomeRoutes(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World Web App")
	render.RenderTemplate(w, "home.page.tmpl")
}

//func HomeRoutes(w http.ResponseWriter, r *http.Request) {
//
//	io.WriteString(w, "<h1>home</h1>")
//
//	fmt.Println("home page ...")
//}

//func addValues(x, y int) int {
//	return x + y
//}
