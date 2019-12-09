package main

import "fmt"

//Modifiez l'exemple pour trop remplir le buffer et voir ce qui se passe.
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
