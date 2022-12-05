package main

import "fmt"

func main() {
	fmt.Println("ready...")
	defer fmt.Println("done")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("go!")
}
