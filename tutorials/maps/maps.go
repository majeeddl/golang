package main

import "fmt"

/*

Go Maps
	Maps are used to store data values in key:value pairs.

	Each element in a map is a key:value pair.

	A map is an unordered and changeable collection that does not allow duplicates.

	The length of a map is the number of its elements. You can find it using the len() function.

	The default value of a map is nil.

	Maps hold references to an underlying hash table.

	Go has multiple ways for creating maps.
*/

func main() {
	fmt.Print("Maps")

	var x = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	y := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", x)
	fmt.Printf("b\t%v\n", y)

	//Creating Maps Using Using make()Function:
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	//Updating and Adding Map Elements
	a["year"] = "1970" // Updating an element
	a["color"] = "red" // Adding an element

	fmt.Println(a)

	//Remove Element from Map

	delete(a, "year")
	fmt.Println(a)

	//Maps Are References
	//Maps are references to hash tables.
	//If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

	//Iterating Over Maps
	for k, v := range a {
		fmt.Printf("%v : %v", k, v)
	}

	//Iterate Over Maps in a Specific Order
	a1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b2 = []string{} // defining the order
	b2 = append(b2, "one", "two", "three", "four")

	for k, v := range a1 { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b2 { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}
}
