package main

func main() {
	cards := newDeck()
	hand, remainingHand := deal(cards, 5)

	// cards.print()
	hand.print()
	remainingHand.print()
}
