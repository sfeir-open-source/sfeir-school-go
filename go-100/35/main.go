package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(250 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		}
	}
}
