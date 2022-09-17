package main

import (
	"fmt"
	"net/http"
	"webapp/routes"

	// "net/http"
	"webapp/config"
	// "webapp/routes"
)

func main() {

	config.SetConfig()
	// http.HandleFunc("/home" , routes.HomeRoutes)

	http.HandleFunc("/", routes.HomeRoutes)
	http.HandleFunc("/about", routes.AboutRoutes)
	http.HandleFunc("/divide", routes.DivideRoutes)

	fmt.Println(config.AppConfig.PORT)

	http.ListenAndServe(":"+config.AppConfig.PORT, nil)

}
