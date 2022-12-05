package main

import "fmt"

type MyStruct struct {
	myInt int
}

// TOOD: créer une méthode map qui prend une liste générique en paramètre ainsi qu'une fonction et retourne la liste modifiée

func main() {
	myStructs := []MyStruct{
		{myInt: 1},
		{myInt: 2},
		{myInt: 3},
		{myInt: 4},
	}
	// Mulitiplier chaque élement du tableau par lui même
	fmt.Print(myStructs)

	myInts := []float64{
		1.2, 2.4, 3.6, 4.8,
	}
	// Garder la partie entière chaque élement du tableau
	fmt.Print(myInts)
}
