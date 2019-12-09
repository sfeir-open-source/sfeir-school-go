package main

import "fmt"

// Aigle est une structure représentant un aigle
type Aigle struct{}

func (a Aigle) Mange() {
	fmt.Println("L'aigle mange.")
}
func (a Aigle) Vole() {
	fmt.Println("L'aigle vole.")
}

// A380 est une structure représentant un Airbus A380
type A380 struct{}

func (a A380) Vole() {
	fmt.Println("L'A380 vole.")
}

// Oiseau est une interface qui définit le comportement d'un oiseau qui mange et vole.
type Oiseau interface {
	Mange()
	Vole()
}

// Avion est une interface qui définit le comportement d'un avion qui vole.
type Avion interface {
	Vole()
}

func main() {
	// Aigle implémente bien l'interface Oiseau
	var oiseau Oiseau = Aigle{}

	// A380 implémente bien l'interface Avion
	var avion Avion = A380{}

	oiseau.Vole()
	oiseau.Mange()
	avion.Vole()

	// Aigle implémente aussi l'interface Avion
	avion = Aigle{}

	avion.Vole()
}
