package main

import (
	"fmt"
)

func Filter[V comparable](collection []V, predicate func(item V, index int) bool) []V {
	result := []V{}

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

type MyStruct struct {
	myNbr int
}

func main() {
	list := []MyStruct{{myNbr: 1}, {myNbr: 2}, {myNbr: 3}, {myNbr: 4}}

	result := Filter(list, func(nbr MyStruct, index int) bool {
		return nbr.myNbr%2 == 0
	})
	fmt.Printf("%v", result)

	listInt := []int{1, 2, 3, 4}

	resultInt := Filter(listInt, func(nbr int, index int) bool {
		return nbr%2 == 0
	})

	fmt.Printf("%v", resultInt)
}
