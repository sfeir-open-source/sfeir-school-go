package main

import "fmt"

func main() {
	a := 3
	b := 5
	if a < b {
		fmt.Printf("%d is lower than %d\n", a, b)
	}

	if c := a*2 - b; c < a {
		fmt.Printf("%d is lower than %d\n", c, a)
	} else {
		fmt.Printf("%d is greater than %d\n", c, a)
	}
}
