package main

import "fmt"

func say(s string, n int, ch chan string) {

	for i := 0; i < n; i++ {
		ch <- s
	}
	ch <- "Goodbye"
	close(ch)
}

func main() {
	ch := make(chan string)
	go say("Bonjour", 3, ch)

	s, ok := <-ch
	fmt.Printf("%q %v\n", s, ok)
	s, ok = <-ch
	fmt.Printf("%q %v\n", s, ok)
	s, ok = <-ch
	fmt.Printf("%q %v\n", s, ok)
	s, ok = <-ch
	fmt.Printf("%q %v\n", s, ok)
	s, ok = <-ch
	fmt.Printf("%q %v\n", s, ok)
}
