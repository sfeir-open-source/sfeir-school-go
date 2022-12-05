package main

import "fmt"

//TODO: créez un type "person" qui contiendra le prénom et le nom d'une personne.

//TODO: créez une méthode 'fullname' sur le type *person qui retourne la concaténation du prénom et du nom.
//La méthode devra gérer le cas où la personne est nil.

func main() {
	jack := &person{"Jack", "Sparrow"}
	fmt.Println(jack.fullname())
	jack = nil
	fmt.Println(jack.fullname())
}
