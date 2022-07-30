package goroutin

import "fmt"

// on buffer channel 
func RunChannels(){

	fmt.Println("*******************")
	fmt.Println("***  start channels ***")

	ch := make(chan int)

	go print(ch)

	//program freez until get data
	message := <- ch

	fmt.Println(message)

	fmt.Println("*******************")
	fmt.Println("***  end channels ***")
}


// buffer channel 
func RunBufferChannels(){

	fmt.Println("*******************")
	fmt.Println("***  start buffer channels ***")

	ch := make(chan int,10)

	go print(ch)

	message := <- ch

	fmt.Println(message)

	fmt.Println("*******************")
	fmt.Println("***  end buffer channels ***")
}


func print(ch chan int){
	ch <- 10
}