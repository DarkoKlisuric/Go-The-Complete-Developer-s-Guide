package main

import (
	"fmt"
	"strings"
)


// Create a new type of 'deck'
// which is a slice of stings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace", "Two",
		"Three", "Four",
		"Five", "Six",
		"Seven", "Eight",
		"Nine", "Ten",
		"Queen", "Jack", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}
// resivers
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
// return two types of deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
//joining a slice to string
func (d deck) toString() string {
	return strings.Join([]string(d) , ",")
}
