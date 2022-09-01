package main

import "fmt"

 
func main() {
  card := newCard();
  fmt.Println(card)

  // cards := deck{ "one" , "two"}

  // cards = append(cards, "three")

  cards := newDeck()

  fmt.Println(cards)

  cards.print()

  // for i,card := range cards{
  //   fmt.Println(i,card)
  // }
}

func newCard() string {
	return "Five of dimonds"
}


