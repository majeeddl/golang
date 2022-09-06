package main

import "fmt"


func main(){

	colors := map[string]string{
		"red" : "red",
	}


	fmt.Println(colors)

	// var colorsOne map[string]string

	colorsOne := make(map[string]string)

	colorsOne["white"] = "white"

	fmt.Println(colorsOne)


	//iterating over maps

	for key, value := range colorsOne{
		fmt.Println(key, value)
	}
}