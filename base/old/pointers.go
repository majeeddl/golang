package pointers

import "fmt"


func CheckPointers(){

    fmt.Println("*******************")
	fmt.Println("***  begin pointers ***")

	a := 1

	increment(a)

	fmt.Println("a normal increment" ,a)

	incrementPointer(&a)

	fmt.Println("a pointer increment" ,a)

	fmt.Println("*******************")
	fmt.Println("***  end  pointers ***")
}


func incrementPointer(a *int){
	*a++
	*a++
	*a++
}

func increment(a int){
	a++
}