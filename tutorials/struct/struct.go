package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}

//A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

//While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

//A struct can be useful for grouping data together to create records.

func main() {
	fmt.Print(" Struct ")

	var person1 Person

	person1.name = "majeed"
	person1.age = 36
	person1.job = "developer"

	fmt.Println("Name: ", person1.name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Job: ", person1.job)
	fmt.Println("Salary: ", person1.salary)

	// Print Pers1 info by calling a function
	printPerson(person1)
}
