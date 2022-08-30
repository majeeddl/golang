package main

import "fmt"

var deckSize int
 
func main() {
  card := newCard();
  fmt.Println(card)
}

func newCard() string {
	return "Five of dimonds"
}


