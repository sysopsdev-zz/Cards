package main

func main() {
	cards := newDeck()
	cards.shuffle()
	currentHand, cardsRemaining := deal(cards, 5)

	currentHand.print()
	cardsRemaining.print()
}
