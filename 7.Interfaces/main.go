package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func (englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}


