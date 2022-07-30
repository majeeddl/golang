package main

import (
	"base/arrays"
	"base/conditions"
	"base/interfaces"
	"base/maps"
	"base/math"
	"base/pointers"
	"base/print"
	"base/structs"
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

	structs.UseStructs()

	pointers.CheckPointers()

	interfaces.RunInterfaces()
}
