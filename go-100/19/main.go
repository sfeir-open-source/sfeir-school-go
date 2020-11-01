package main

import (
	"fmt"
	"math"
)

// déclaration d'un nouveau type de fonction
type NumberFunc func (float64, float64) float64

//compute prend une fonction en paramètre (qui prend deux float64 en paramètre et retourne un float64)
func compute(fn NumberFunc) float64 {
	return fn(3, 4)
}

//added retourne une closure, c'est à dire une méthode liée à une variable en dehors de son scope.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	// déclaration et assignation de fonction
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// utilisation de la fonctiona déclarée
	fmt.Println(hypot(5, 12))

	// utilisation de compute
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
