
package main
import "fmt"


func main()  {
	
	fmt.Print(" ARRAY " , "\n")


	/*
	var array_name = [length]datatype{values} // here length is defined
	or
	var array_name = [...]datatype{values} // here length is inferred
	*/

	//Note: The length specifies the number of elements to store in the array. 
	//In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).

	var arr1 = [3]int{1,2,3}
	arr2 := [6]int{1,2,3,4,5,6}

	fmt.Println(arr1,arr2)


	//This example declares two arrays (arr3 and arr4) with inferred lengths
	var arr3 = [...]int{1,2,3}
	arr4 := [...]int{2,3,4,5}

	fmt.Print(arr3,"\n",arr4)


	//Access Elements of an Array
	prices := [3]int{10,20,30}
	fmt.Println(prices[2])

	//Change Elements of an Array
	prices[2] = 40
	fmt.Println(prices[2])

	//Initialize Only Specific Elements

	arr5 := [5]int{1:10,4:30}
	fmt.Println(arr5)


	//Find the Length of an Array
	fmt.Print(len(arr5))
}