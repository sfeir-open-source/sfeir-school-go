package main

import "fmt"

func multiplicateurPar(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {
	mult := multiplicateurPar(2)
	for i := 0; i < 10; i++ {
		fmt.Println(mult(i))
	}
}
