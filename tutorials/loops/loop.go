package main

import "fmt"

func main() {

	fmt.Printf(" Loop \n")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	//The continue Statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	//The break Statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}


	//The Range Keyword
	fruits := [3]string{"apple","banana","orange"}

	for idx,val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

}
