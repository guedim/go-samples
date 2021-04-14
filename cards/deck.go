package main

import (
	"fmt"
	"strings"
)

// Create a new type of deck
// wich is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "C"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
