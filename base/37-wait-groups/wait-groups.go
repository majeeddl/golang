package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("worker", id, "started job")
	time.Sleep(time.Second)
	fmt.Println("worker", id, "finished job")
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
