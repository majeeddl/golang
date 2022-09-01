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

  one, two := deal(cards,1)

  one.print()
  two.print()
  // for i,card := range cards{
  //   fmt.Println(i,card)
  // }

  greeting := "Hi there!"
  fmt.Println([]byte(greeting))

  fmt.Println(cards.toString())
}

func newCard() string {
	return "Five of dimonds"
}


