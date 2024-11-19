package main

import "fmt"

func main() {
	fmt.Print("Slices", "\n")

	/*

		Slices are similar to arrays, but are more powerful and flexible.

		Like arrays, slices are also used to store multiple values of the same type in a single variable.

		However, unlike arrays, the length of a slice can grow and shrink as you see fit.

		In Go, there are several ways to create a slice:

			Using the []datatype{values} format
			Create a slice from an array
			Using the make() function

	*/

	//Create a Slice With []datatype{values}
	mySlice := []int{}
	mySlice = []int{1, 2, 3}

	fmt.Print(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	//Create a Slice From an Array
	var myArray = [3]int{1, 2, 3}
	mySlice = myArray[1:3]

	fmt.Print(mySlice, "\n")

	//Create a Slice With The make() Function
	//Note: If the capacity parameter is not defined, it will be equal to length.

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	//Change Elements of a Slice
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Print("\n")
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	//Append Elements To a Slice
	prices = append(prices, 40, 50)
	fmt.Print("\n")
	fmt.Println(prices)

	//Append One Slice To Another Slice
	myslice3 := []int{1, 2, 3}
	myslice4 := []int{4, 5, 6}
	myslice5 := append(myslice3, myslice4...)
	fmt.Print("\n")
	fmt.Println(myslice5)

	/*
			Memory Efficiency
		 		When using slices, Go loads all the underlying elements into the memory.

				If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

				The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.

	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
