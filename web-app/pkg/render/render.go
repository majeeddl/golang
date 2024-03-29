package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate /Render template  renders templates using html/tmpl
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	fmt.Println("./templates/" + tmpl)
	//, "./templates/base.layout.tmpl"
	parsedTmp, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parsedTmp.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template :", err)
		return
	}

}
