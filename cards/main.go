package main

import "fmt"

func main() {
	card := newCard()

	fmt.Print(card)
}

func newCard() string {
	return "Five of Diamonds"
}
