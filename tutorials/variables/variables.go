package main

import "fmt"

func main() {

	/*
		Declaring (Creating) Variables
	*/

	/*
		1. With the var keyword
		Note: You always have to specify either type or value (or both).
	*/
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	/*
		2. With the := sign:
		Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).
		Note: It is not possible to declare a variable using :=, without assigning a value to it.
		Note: Can only be used inside functions
	*/
	x := 2 //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	/*
		Variable Declaration Without Initial Value
		In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:
	*/

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	/*
		Value Assignment After Declaration
	*/
	var studentNew string
	studentNew = "John"
	fmt.Println(studentNew)



	// Go Multiple variable declaration

	var e,f,g int = 1,2,3;

	fmt.Println(e,f,g)

	//Note: If you use the type keyword, it is only possible to declare one type of variable per line.

	var a1,a2 = 6,"Hello"

	fmt.Println(a1,a2)


	//Go Variable Declaration in a Block

	var (
		a11 int
		b11 int = 1
		c11 string = "hello"
	)

	fmt.Println(a11,b11,c11)


	/*
	Go variable naming rules:

		A variable name must start with a letter or an underscore character (_)
		A variable name cannot start with a digit
		A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
		Variable names are case-sensitive (age, Age and AGE are three different variables)
		There is no limit on the length of the variable name
		A variable name cannot contain spaces
		The variable name cannot be any Go keywords

		*/
}
