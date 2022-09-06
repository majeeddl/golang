package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo // is same with contactInfo
}

func main() {

	newPerson := person{firstName: "majeed", lastName: "d", contactInfo: contactInfo{email: "majeed.dl@gmail.com", zipCode: 12}}

	fmt.Println(newPerson)

	var newPersonTwo person

	fmt.Printf("%+v", newPersonTwo)

	newPersonTwo.firstName = "Ali"
	newPersonTwo.lastName = "d"
	newPersonTwo.contactInfo = contactInfo{email: "ali@gmail.com", zipCode: 12}

	fmt.Println(newPersonTwo)

	newPersonTwo.print()

	// newPersonTwoPointer := &newPersonTwo;
	newPersonTwo.updateName("John")

	newPersonTwo.print()

	name := "Bill"
 
    fmt.Println(*&name)

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
