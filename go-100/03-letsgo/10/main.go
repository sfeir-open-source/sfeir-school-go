package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("C'est bientôt le weekend ?")
	today := time.Now().Weekday()
	switch today {
	case time.Saturday:
		fallthrough
	case time.Sunday:
		fmt.Println("C'est le weekend !")
	case time.Friday:
		fmt.Println("Bientôt, mais pas encore.")
	default:
		fmt.Println("Dans trop longtemps.")
	}
}
