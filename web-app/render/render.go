package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	fmt.Println("./templates/" + tmpl)
	parsedTmp, _ := template.ParseFiles("./templates/" + tmpl, "./templates/baase.layout.tmpl")

	err := parsedTmp.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template :", err)
		return
	}

}
