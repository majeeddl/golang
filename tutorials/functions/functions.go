package main

import "fmt"

func familyName(fname string, age int) string {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
	return fname
}

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

func multipleReturn(x int, y string) (result1 int, txt1 string) {
	result1 = x + x
	txt1 = y + " World!"
	return
}



//Go Recursion Functions
//Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.
func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}



func main() {
	fmt.Printf(" Functions \n")

	result := familyName("John", 32)

	fmt.Printf("\n %v", result)

	//Named Return Values
	fmt.Println(myFunction(1, 2))

	// Multiple return values
	fmt.Println(multipleReturn(12, "Hello"))

	a, b := multipleReturn(5, "Hello")
	fmt.Println(a, b)

	_, b1 := multipleReturn(5, "Hello")
	fmt.Println(b1)


	testcount(1)
}
