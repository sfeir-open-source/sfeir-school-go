package main

import (
	"fmt"
	"net/http"
	"time"
)

const url = "https://www.google.fr/search?q=golang"

func curl() string {
	resp, err := http.Get(url)
	if err == nil {
		return resp.Status
	}
	return "Error"
}

func main() {
	defer trackTimeElapsed(time.Now())

	const reqCount = 20
	for i := 0; i < reqCount; i++ {
		result := curl()
		fmt.Printf(url+" responded with HTTP status %s\n", result)
	}
}

func trackTimeElapsed(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)
}
