package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()
	//hand, remainingDeck := deal(cards, 5)

	//hand.print()
	//remainingDeck.print()

	fmt.Println("-------")

	fmt.Println(cards.toString())

}
