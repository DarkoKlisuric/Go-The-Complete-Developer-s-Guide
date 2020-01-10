package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://app.e-udruge.eu/",
		"http://dora.ci-sdz.hr/",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	//receive value from channel
/*

	for {
		go checkLink(<-c, c)
	}
*/

	for l := range c{
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		//sending a message into channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}