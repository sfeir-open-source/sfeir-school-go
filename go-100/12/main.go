package main

import (
	"errors"
	"fmt"
)

//Distance en mètres
type Distance float64

//Time en secondes
type Time float64

//Speed en m/s
type Speed float64

func computeSpeed(d Distance, t Time) (Speed, error) {
	if t == 0 {
		return 0, errors.New("impossible de calculer une vitesse sur un temps nul")
	}
	return Speed(float64(d) / float64(t)), nil
}

func main() {
	// _ permet d'ignorer la seconde valeur retournée (ici, une erreur)
	speed, _ := computeSpeed(130000, 3600)

	fmt.Printf("Vous allez à %fm/s\n", speed)
}
