package main

import (
	"errors"
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

func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {

	//x := 100.0
	//y := 10.0
	f, err := divideValues(100.0, 10.0)

	if err != nil {
		fmt.Fprintf(w, "Can not divide by 0")
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("can not divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

func main() {

	config.SetConfig()
	// http.HandleFunc("/home" , routes.HomeRoutes)

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(config.AppConfig.PORT)

	http.ListenAndServe(":"+config.AppConfig.PORT, nil)

}
