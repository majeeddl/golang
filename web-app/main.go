package main

import (
	"fmt"
	"net/http"
	"webapp/routes"
)

func main(){

	http.HandleFunc("/home" , routes.HomeRoutes)

	fmt.Println("server is starting on port 3000 ...")

	http.ListenAndServe(":3000",nil)

}