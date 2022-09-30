package main

import (
	"fmt"
	"net/http"
	"webapp/pkg/handlers"
	// "net/http"
	"webapp/config"
	// "webapp/routes"
)

func main() {

	config.SetConfig()
	// http.HandleFunc("/home" , routes.HomeRoutes)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	//http.HandleFunc("/divide", routes.DivideRoutes)

	fmt.Println(config.AppConfig.PORT)

	http.ListenAndServe(":"+config.AppConfig.PORT, nil)

}
