package routes

import (
	"fmt"
	"net/http"
)

func HomeRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Web App")
}

//func HomeRoutes(w http.ResponseWriter, r *http.Request) {
//
//	io.WriteString(w, "<h1>home</h1>")
//
//	fmt.Println("home page ...")
//}

func addValues(x, y int) int {
	return x + y
}
