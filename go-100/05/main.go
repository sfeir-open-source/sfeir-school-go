package main

import "fmt"

var a bool = false // déclaration inutilement complète : var name type = expression
var b = false      // le type est inféré par la valeur initiale
var c bool         // false par défaut (la "valeur-zéro" d'un booléen est false)

func uneFonction() {
	var w bool = false // déclaration inutilement complète : var name type = expression
	var x = false      // le type est inféré par la valeur initiale
	var y bool         // la valeur initiale est la "valeur-zéro" du type de la vairable (la "valeur-zéro" d'un booléen est false)
	z := false         // déclaration courte : à l'intérieur d'une fonction on peut utiliser := au lieu de var et le type est inféré par la valeur initiale

	fmt.Printf("w = %v and its type is %[1]T\n", w)
	fmt.Printf("x = %v and its type is %[1]T\n", x)
	fmt.Printf("y = %v and its type is %[1]T\n", y)
	fmt.Printf("z = %v and its type is %[1]T\n", z)
}

func main() {
	fmt.Printf("a = %v and its type is %[1]T\n", a)
	fmt.Printf("b = %v and its type is %[1]T\n", b)
	fmt.Printf("c = %v and its type is %[1]T\n", c)
	uneFonction()
}
