package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
			cards = append(cards, value+" of "+suit)
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
//Writing file on hard drive
func (d deck) saveToFile(filename string) error {
	//converting string to byte slice
	byteSlice := []byte(d.toString())
	return ioutil.WriteFile(filename, byteSlice, 0666)
}
//Reading deck from hard drive
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//log the error and entirely quit the program
		fmt.Println("Error:", err)
		//if error is not nil , print error
		os.Exit(1)
	}
	// return slice of strings from byte slice
	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	//return datetime now in intiger
	t := time.Now().UnixNano()
	//seed value of rand
	source := rand.NewSource(t)
	//create new source of random number
	r := rand.New(source)

	for i := range d {
		// generate random number between 0 and lenght of slice
		newPostion := r.Intn(len(d)-1)
		// swap the elements
		d[i], d[newPostion] = d[newPostion], d[i]
	}
}