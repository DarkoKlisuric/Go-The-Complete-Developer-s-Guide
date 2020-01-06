package main

func main()  {
	cards := newDeck()

	/*
	hand, remainingCards := deal(cards, 12)

	hand.print()
	remainingCards.print()
	*/
	/*
	fmt.Println(cards.toString())
	*/
	cards.saveToFile("FirstDeck")
}
