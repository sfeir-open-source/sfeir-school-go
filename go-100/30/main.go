package main

import (
	"fmt"
	"time"
)

func say(s string, ch chan string) {
	fmt.Println("ch <- s attend que quelqu'un lise le canal ch.")
	ch <- s
	fmt.Println("s a été envoyée sur ch.")
}

func main() {
	ch := make(chan string)
	go say("Bonjour", ch)

	time.Sleep(10 * time.Millisecond)
	fmt.Println("s := <-ch attend que quelqu'un envoie des données sur le canal ch.")

	s := <-ch

	time.Sleep(10 * time.Millisecond)
	fmt.Printf("La string '%s' a été lue depuis le canal ch.\n", s)
}
