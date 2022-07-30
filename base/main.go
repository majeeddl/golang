package main

import (
	"base/math"
	"base/print"
	"base/arrays"
	"base/maps"
	"base/conditions"
	"fmt"
)

const hello_world = "Hello World"

func main() {
	fmt.Println(hello_world)

	fmt.Println(print.PrintText("Hello world"))

	fmt.Println(mathLib.Sum(1, 2))

	fmt.Println(mathLib.Multiply(2, 4))

	arrays.PrintArrays()

	maps.PrintMaps()

	conditions.PrintConditions()
}
