package main

import "fmt"

//indexOf prend en paramètre une chaine recherchée et un slice de chaines,
//et retourne l'indice de la chaine recherchée dans le slice, ou -1 si elle n'a pas été trouvée.
func indexOf(search string, s []string) int {
	for i, v := range s {
		if v == search {
			return i
		}
	}
	return -1
}

//split prend en paramètre une chaine recherchée et un slice de chaines,
//et retourne un slice de slice.
//Ex: split([]string{"|", "B", "|", "C", "D", "E", "|", "F", "|"}) = [[B] [C D E] [F]]
func split(search string, s []string) [][]string {
	splitted := make([][]string, 0, 5)
	lastIndex := 0
	for i, v := range s {
		if v == search {
			if i > 0 {
				splitted = append(splitted, s[lastIndex:i])
			}
			lastIndex = i + 1
		}
	}
	return splitted
}

func main() {
	var jours = []string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
	var index int

	index = indexOf("lundi", jours)
	fmt.Printf("Le lundi est le %der jour de la semaine\n", index+1)

	index = indexOf("jeudi", jours)
	fmt.Printf("Le jeudi est le %dème jour de la semaine\n", index+1)

	index = indexOf("plop", jours)
	fmt.Printf("Plop est un jour de la semaine : %v\n", index > -1)

	var toBeSplitted = []string{"|", "B", "|", "C", "D", "E", "|", "F", "|"}
	splitted := split("|", toBeSplitted)
	fmt.Println(splitted)
}
