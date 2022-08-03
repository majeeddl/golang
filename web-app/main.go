package main

import (
	"fmt"
	// "net/http"
	"webapp/config"
	// "webapp/routes"
)

func main(){

	config.SetConfig()
	// http.HandleFunc("/home" , routes.HomeRoutes)

	// fmt.Println("server is starting on port 3000 ...")

	// http.ListenAndServe(":3000",nil)
	fmt.Println(config.AppConfig.PORT)
}