package main

import (
	"fmt"
	"net/http"
	"time"
)

const url = "https://www.google.fr/search?q=golang"

func curl(c chan string) {
	resp, err := http.Get(url)
	if err == nil {
		c <- resp.Status
	} else {
		c <- "Error"
	}
}

func main() {
	defer trackTimeElapsed(time.Now())

	c := make(chan string)

	const reqCount = 20
	for i := 0; i < reqCount; i++ {
		go curl(c)
	}
	for i := 0; i < reqCount; i++ {
		result := <-c
		fmt.Printf(url+" responded with HTTP status %s\n", result)
	}
}

func trackTimeElapsed(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)
}
