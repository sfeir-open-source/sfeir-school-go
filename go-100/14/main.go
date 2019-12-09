package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // pointe vers i
	fmt.Println(*p) // lit i via le pointeur
	*p = 21         // modifie i via le pointeur
	fmt.Println(i)  // affiche la nouvelle valeur de i

	p = &j         // pointe vers j
	*p = *p / 37   // divise j via le pointeur
	fmt.Println(j) // affiche la nouvelle valeur de j
}
