package main

import "fmt"

func main() {
	fmt.Printf("switch \n")

	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Not a weekday")
	}
}
