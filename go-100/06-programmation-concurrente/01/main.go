package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("Je parle quand je veux !")
	say("Arrête de parler en même temps que moi !")
}
