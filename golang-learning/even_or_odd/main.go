package main

import "fmt"


func main(){

	slice := [10]int{}

	for _,num := range slice {
		if num % 2 == 0 {
			fmt.Println(num, " is even")
		}else{
			fmt.Println(num, " is odd")
		}
	}
}