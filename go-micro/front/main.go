package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<h1>hello world</h1>")
	})

	err := http.ListenAndServe(":9090", nil)

	fmt.Println("server is starting on port 9090 ...")

	if err != nil {
		log.Panic(err)
	}

}
