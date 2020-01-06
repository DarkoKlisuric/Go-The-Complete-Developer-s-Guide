package main

func main()  {
	/*
	cards := newDeck()
	#1
	hand, remainingCards := deal(cards, 12)
	hand.print()
	remainingCards.print()
	*/
	/*
	#2
	fmt.Println(cards.toString())
	*/
	/*
	#3
	cards.saveToFile("FirstDeck")
	*/
	cards := newDeckFromFile("FirstDeck")
	cards.print()
}
