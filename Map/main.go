package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func test () {
	// all keys are string and all the values are string
	// 1. type
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00FF00",
	}
	fmt.Println(colors)
	// 2. type
	var color map[string]string
	//empty map
	fmt.Println(color)

	// 3. type
	moreColors := make(map[string]string)
	moreColors["white"] = "#fffff"

	delete(moreColors, "white")

	fmt.Println(moreColors)
}