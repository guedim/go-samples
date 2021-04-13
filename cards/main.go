package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("---------")

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	fmt.Println("*********")
	remainingDeck.print()

}
