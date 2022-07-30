package routes

import (
	"fmt"
	"io"
	"net/http"
)

func HomeRoutes(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "<h1>home</h1>")

	fmt.Println("home page ...")
}
