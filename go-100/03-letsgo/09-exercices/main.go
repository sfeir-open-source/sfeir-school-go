package main

import "fmt"

//TODO: écrire une fonction qui prend en paramètre un entier et un nombre max d'itérations à effectuer, et qui effectue l'algorithme de
//la conjecture de Syracuse (voir README.md).
//Elle devra retourner un booléen indiquant si le nombre 1 a été atteint, suivi du nombre d'itérations qui ont été effectuées.

func main() {
	const maxIterationCount = 1000
	const integer = 98375
	reached, iterationCount := syracuse(integer, maxIterationCount)
	if reached {
		fmt.Printf("L'algorithme a permis d'atteindre le nombre 1 à partir du nombre %d en %d itérations\n", integer, iterationCount)
	} else {
		fmt.Printf("L'algorithme n'a pas permis d'atteindre le nombre 1 à partir du nombre %d en moins de %d itérations\n", integer, maxIterationCount)
	}
}
