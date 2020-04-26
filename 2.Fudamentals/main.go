package main

import (
	"fmt"
)

func main() {
	//slices
	cards := []string{"Ace of diamons", newCard()}
	//add new element to slice
	cards = append(cards, "Six of spades")

	//looping with index key
	for i, card := range cards {
		fmt.Println(i, card)
	}
	// looping without index key
	for _, card := range cards {
		fmt.Println(card)
	}
	fmt.Println(cards)

	fmt.Println(newCard())
}

func newCard() string {
	return "Five of diamonds"
}
