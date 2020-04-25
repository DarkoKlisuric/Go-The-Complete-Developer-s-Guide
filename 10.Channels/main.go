package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"9.http://app.e-udruge.eu/",
		"9.http://dora.ci-sdz.hr/",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
/*
	//receive value from channel

	for {
		go checkLink(<-c, c)
	}


	for l := range c {
		go checkLink(l, c)
	}
*/
 	for l := range c {
 		go func(link string) {
			//pause of 3 seconds
 			time.Sleep(3 * time.Second)
 			checkLink(link, c)
		}(l)
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