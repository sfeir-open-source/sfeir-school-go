package main

import (
	"fmt"
)

func fibonacci(iter int) int {
	n, prev := 1, 0
	for i := 0; i < iter; i++ {
		n, prev = n+prev, n
	}
	return n
}

func main() {
	age, nom, enVie := 30, "John", true
	fmt.Println(age, nom, enVie)

	fmt.Println(fibonacci(20))
}
