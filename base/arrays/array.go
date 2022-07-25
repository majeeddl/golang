package arrays

import "fmt"

func PrintArrays() {
	fmt.Println("*******************")
	fmt.Println("***  begin arrays ***")

	a := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)

	fmt.Println(">>> set item in array in index 2 to 22")
	a[1] = 22

	fmt.Println(a)

	fmt.Println("***  end arrays ***")
	fmt.Println("*******************")
}
