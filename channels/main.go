package main

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
	// http.Get() is a blocking call
	// go checkLink(link) will create a new go routine
	// go routine is a lightweight thread
	// main routine will not wait for the go routine to finish
	// main routine will continue to run
	// go routine will run in parallel
	// go routine will not share memory with other go routines
	// go routine will share memory with main routine
}
