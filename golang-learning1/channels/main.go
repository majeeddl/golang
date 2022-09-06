package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "This site might be down")
		c <- "This site might be down"
		return
	}

	fmt.Println(link, " is up")
	c <- "Yep is up"

}
