package main

import "fmt"

func main()  {
	//creating new deck
	cards := newDeck()
	// shuffle and deal 12 cards
	cards.shuffle()
	hand, remainingCards := deal(cards, 12)
	fmt.Println("Your cards:")
	hand.print()
	fmt.Println("Remaining cards:")
	remainingCards.print()
	// saving current combination to file
	cards.toString()
	cards.saveToFile("ShuffleDeck")
	// reading current combination from file
	readDeck := newDeckFromFile("ShuffleDeck")
	readDeck.print()
}
