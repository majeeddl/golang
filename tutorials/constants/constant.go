package main

import "fmt"

func main() {

	const PI = 3.14

	fmt.Println(PI)

	/*

		Constant Rules
			Constant names follow the same naming rules as variables
			Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
			Constants can be declared both inside and outside of a function

	*/

	/*
		Constant Types
			There are two types of constants:

				Typed constants
				Untyped constants


	*/

	//Typed constants
	const A int = 1

	//Untypes constants
	const B = 1

	//Multiple declaration

	const (
		A1 int = 1
		B2     = 3.14
		C2     = "Hi!"
	)

	fmt.Printf("A has value: %v and type: %T\n", A, A)
  	fmt.Printf("B has value: %v and type: %T", C2, C2)
}
