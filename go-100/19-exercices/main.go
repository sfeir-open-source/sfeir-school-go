package main

import "fmt"

func multiplicateurPar(x int) func(int) int {
	// TODO: implémenter cette fonction
}

func main() {
	mult := multiplicateurPar(2)
	for i := 0; i < 10; i++ {
		fmt.Println(mult(i))
	}
}
