package main

import "fmt"

func multiplicateurPar(x int) func(int) int {
	// TODO: implémenter cette fonction
	return nil
}

func main() {

	doubleur := multiplicateurPar(2)
	tripleur := multiplicateurPar(3)

	for i := 0; i < 10; i++ {
		fmt.Println(doubleur(i))
		fmt.Println(tripleur(i))
	}
}
