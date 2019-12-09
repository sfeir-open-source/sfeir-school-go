//Imports factorisés
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Now you have %d problem(s). \n", rand.Intn(10))
}
