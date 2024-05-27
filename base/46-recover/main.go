package main

import "fmt"

func myPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()

	myPanic()

	/*
	This code will not run, because mayPanic panics. The execution of main stops at the point of the panic and resumes in the deferred closure.
	*/
	fmt.Println("This will not be printed")
}
