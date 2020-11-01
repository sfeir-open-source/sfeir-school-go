package main

import (
	"fmt"
)

func count(n int, c chan int) {
	x := 0
	for i := 0; i < n; i++ {
		c <- x
		x++
	}
	close(c)
}

func main() {
	c := make(chan int, 2)
	go count(10, c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Le channel a été fermé : on sort de la boucle for.")
}
