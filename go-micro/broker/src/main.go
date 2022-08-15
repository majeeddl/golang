package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "9091"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s \n", port)

	//define http server
	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}

	//start the server
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("broker")
}
