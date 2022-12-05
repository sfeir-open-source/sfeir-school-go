package main

import "fmt"

type MonInterface interface {
	M()
}

type MaStructure struct {
	S string
}

// Le type *MaStructure implémente I
func (t *MaStructure) M() {
	fmt.Println(t.S)
}

func main() {
	t := MaStructure{"hello"}
	maFonction(t) // erreur de compilation : le type MaStructure n'implémente pas MonInterface.
}

func maFonction(param MonInterface) {
	fmt.Printf("(%v, %T)\n", param, param)
}
