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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")

	// http.Get() is a blocking call
	// go checkLink(link) will create a new go routine
	// go routine is a lightweight thread
	// main routine will not wait for the go routine to finish
	// main routine will continue to run
	// go routine will run in parallel
	// go routine will not share memory with other go routines
	// go routine will share memory with main routine
}
