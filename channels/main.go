package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think!"
		return
	}

	fmt.Println(link, "is up!")

	c <- "Yep it's up!"

	// http.Get() is a blocking call
	// go checkLink(link) will create a new go routine
	// go routine is a lightweight thread
	// main routine will not wait for the go routine to finish
	// main routine will continue to run
	// go routine will run in parallel
	// go routine will not share memory with other go routines
	// go routine will share memory with main routine
}
