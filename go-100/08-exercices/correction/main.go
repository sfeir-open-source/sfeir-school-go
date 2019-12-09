package main

import "fmt"

//inverse prend un nombre décimal en paramètre et retourne son inverse. Si le nombre est nul, on retourne false (seconde valeur de retour)
func inverse(x float64) (float64, bool) {
	if x == 0 {
		return 0, false
	}
	return 1 / x, true
}

func main() {
	// Calcul de l'inverse de 3
	if inv, ok := inverse(3); ok {
		fmt.Printf("L'inverse de %f est %f\n", 3.0, inv)
	} else {
		fmt.Printf("Impossible de calculer l'inverse de %f", 3.0)
	}

	// Calcul de l'inverse de 0 !
	if inv, ok := inverse(0); ok {
		fmt.Printf("L'inverse de %f est %f\n", 0.0, inv)
	} else {
		fmt.Printf("Impossible de calculer l'inverse de %f\n", 0.0)
	}
}
