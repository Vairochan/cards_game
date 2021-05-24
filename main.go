package main

func main() {

	// Cards := newDeck()

	// Cards.print()

	// hand, remainingDeck := deal(Cards, 5)
	// // fmt.Println(hand, remainingDeck)

	// hand.print()
	// remainingDeck.print()

	// // newString := hand.toStringByInbuild()

	// // fmt.Println(newString)
	// Cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	// cards.print()

	card := newDeck()
	card.shuffle()
	card.print()

}
