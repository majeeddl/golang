package main

import (
	"fmt"
	"net/http"

	// "net/http"
	"webapp/config"
	// "webapp/routes"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Web App")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page Web App")
}

func main() {

	config.SetConfig()
	// http.HandleFunc("/home" , routes.HomeRoutes)

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)

	fmt.Println(config.AppConfig.PORT)

	http.ListenAndServe(":"+config.AppConfig.PORT, nil)

}
