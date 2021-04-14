package main

var fileName string

func main() {

	fileName = "my_cards.txt"

	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	cards.saveToFile(fileName)

	cardsFromFile := readFromFile(fileName)
	cardsFromFile.print()

	cards.suffle()
	cards.print()
}
