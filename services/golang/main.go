package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	// hand, remaningDeck := deal(cards, 5)

	// hand.print()
	// remaningDeck.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
}
