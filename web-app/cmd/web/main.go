package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"webapp/pkg/config"
	"webapp/pkg/databases"
	"webapp/pkg/models"

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

	db := databases.DB()

	user := models.User{Name: "Jinzhu", Age: 18}

	result := db.Create(&user)

	fmt.Println(result)

	srv := &http.Server{
		Addr:    ":" + config.AppConfig.PORT,
		Handler: routes(),
	}

	err := srv.ListenAndServe()

	log.Fatal(err)

}

func run(param int) error {
	fmt.Println("RUN")

	if param > 2 {
		return nil
	}

	return errors.New("param in less than 2")
}
