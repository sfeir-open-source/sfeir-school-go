package main

import "fmt"

var jours = []string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

func main() {
	for i, v := range jours {
		fmt.Printf("jours[%d] = %s\n", i, v)
	}

	puissancesDe2 := make([]int, 10)
	for i := range puissancesDe2 {
		puissancesDe2[i] = 1 << uint(i)
	}
	for _, value := range puissancesDe2 {
		fmt.Printf("%d\n", value)
	}
}
