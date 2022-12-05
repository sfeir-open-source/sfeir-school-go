package main

import (
	"fmt"
	"math"
)

type MyStruct struct {
	myInt int
}

type number interface {
	int | float64
}

// TOOD: créer une méthode map qui prend une liste générique en paramètre ainsi qu'une fonction et retourne la liste modifiée
func Map[T any, R any](collection []T, apply func(item T, index int) R) []R {
	var result []R
	// result := make([]T, len(collection))
	for i, item := range collection {
		result = append(result, apply(item, i))
		//result[i] = apply(item, i)
	}
	return result
}

func main() {
	myStructs := []MyStruct{
		{myInt: 1},
		{myInt: 2},
		{myInt: 3},
		{myInt: 4},
	}
	// Mulitiplier chaque élement du tableau par lui même
	StructsSquared := Map(myStructs, func(item MyStruct, index int) int {
		return item.myInt * item.myInt
	})
	fmt.Println("New element of struct")
	fmt.Print(StructsSquared)

	myFloats := []float64{
		1.2, 2.4, 3.6, 4.8, 10.5,
	}
	// Garder la partie entière chaque élement du tableau
	myFloatsFloored := Map(myFloats, func(item float64, index int) float64 {
		return math.Floor(item)
	})
	fmt.Println("\n Floor of Ints")
	fmt.Print(myFloatsFloored)
}
