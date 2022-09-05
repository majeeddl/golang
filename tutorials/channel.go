package main

import (
	"fmt"
	"math/rand"
)


const numPool = 10


func CalculateValue(intChan chan int){
  randomNumber := rand.Intn(numPool)
  intChan <- randomNumber
}


func main(){
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <- intChan
	fmt.Println(num)
}