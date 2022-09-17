package routes

import (
	"errors"
	"fmt"
	"net/http"
)

func DivideRoutes(w http.ResponseWriter, r *http.Request) {

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
