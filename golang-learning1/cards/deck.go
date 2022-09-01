package main

import "fmt"

// create a new type of 'deck'
// which is a slice of string
type deck []string

func newDeck() deck{
	cards := deck{}
	
	cardSuits := []string{ "Spades","Dimonds"}

	cardValues := []string{ "Ace","Two"}

	for _,suit := range cardSuits{
		for _,value := range cardValues {
			cards = append(cards, suit+ " of " + value)
		}
	}

	return cards
}

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
}


