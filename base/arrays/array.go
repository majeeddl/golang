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

	fmt.Println(">>> print a string array")
	b := [3]string { "ali" , "john" , "majeed"}

	fmt.Println(b)

	fmt.Println(">>> slices ")

	c := []int { 10, 20}

	fmt.Println(c)
	fmt.Println("len of c =>" , (len(c)))

	c = append(c, 1000,2000)

	fmt.Println(c)
	fmt.Println("len of c =>" , (len(c)))

	index := 2

	c = append(c[:index],c[index+1:]...)

	fmt.Println(c)
	fmt.Println("len of c =>" , (len(c)))

	fmt.Println("***  end arrays ***")
	fmt.Println("*******************")
}
