package main

import (
	"fmt"
	"testing"
)


func TestNewDeck(t *testing.T){
	d := newDeck()
	fmt.Println(d)

	if len(d) != 4 {
		t.Errorf(" expected deck length of 4, but got %s", string(rune(len(d))))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf(" first element in deck is not correct")
	}
}