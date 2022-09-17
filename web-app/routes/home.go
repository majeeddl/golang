package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeRoutes(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World Web App")
	renderTemplate(w, "home.page.tmpl")
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

func renderTemplate(w http.ResponseWriter, tmpl string) {
	fmt.Println("./templates/" + tmpl)
	parsedTmp, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTmp.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template :", err)
		return
	}

}
