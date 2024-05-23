package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
