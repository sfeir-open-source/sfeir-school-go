package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Scale utilise un récepteur de pointeur car il doit modifier la valeur de v.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Abs utilise un récepteur de valeur. Dans les faits, on utiliserait plutôt
//un récepteur de pointeur même si Abs n'a pas besoin de modifier la valeur de v,
//car cela éviterait de créer une copie de v à chaque appel de la méthode.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	pv := &v
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	pv.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
