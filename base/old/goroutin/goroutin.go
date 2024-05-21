package goroutin

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func RunGoroutin() {

	fmt.Println("*******************")
	fmt.Println("***  start goroutin ***")

	waitGroup.Add(1)
	go task1()

	waitGroup.Add(1)
	go task2()

	waitGroup.Add(1)
	go task3()

	waitGroup.Wait()

	fmt.Println("*******************")
	fmt.Println("***  end goroutin ***")
}

func task1() {
	for i := 0; i < 10; i++ {
		fmt.Println("task1 : ", i)
	}

	waitGroup.Done()
}

func task2() {
	for i := 11; i < 20; i++ {
		fmt.Println("task2 : ", i)
	}

	waitGroup.Done()
}

func task3() {

	for i := 21; i <= 30; i++ {
		fmt.Println("task3 : ", i)
	}

	waitGroup.Done()

}
