package main

import (
	"fmt"
	"math"
)

//compute prend une fonction en paramètre (qui prend deux float64 en paramètre et retourne un float64)
func compute(fn func(float64, float64) float64) float64 {
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
	// Valeurs de fonction
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
