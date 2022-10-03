package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/pkg/config"
	// "net/http"
	// "webapp/routes"
)

func main() {

	config.SetConfig()
	// http.HandleFunc("/home" , routes.HomeRoutes)

	//http.HandleFunc("/", handlers.Home)
	//http.HandleFunc("/about", handlers.About)
	//http.HandleFunc("/divide", routes.DivideRoutes)

	fmt.Println(config.AppConfig.PORT)

	//http.ListenAndServe(":"+config.AppConfig.PORT, nil)

	srv := &http.Server{
		Addr:    ":" + config.AppConfig.PORT,
		Handler: routes(),
	}

	err := srv.ListenAndServe()

	log.Fatal(err)

}
