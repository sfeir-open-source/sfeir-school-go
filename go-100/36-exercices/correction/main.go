package main

import (
	"fmt"
	"net/http"
	"time"
)

const url = "https://www.google.fr/search?q=golang"

func curl(c chan string, cerr chan error) {
	resp, err := http.Get(url)
	if err == nil {
		c <- resp.Status
	} else {
		cerr <- err
	}
}

func main() {
	defer trackTimeElapsed(time.Now())

	c := make(chan string)
	cerr := make(chan error)

	const reqCount = 100
	for i := 0; i < reqCount; i++ {
		go curl(c, cerr)
	}

	for i := 0; i < reqCount; i++ {
		select {
		case result := <-c:
			fmt.Printf(url+" responded with HTTP status %s\n", result)
		case err := <-cerr:
			fmt.Printf(url+" ERROR: %s\n", err)
		}
	}
}

func trackTimeElapsed(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)
}
