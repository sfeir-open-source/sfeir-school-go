package main

import (
	"fmt"
)

type MyFloat float64

func (pf *MyFloat) IsPositive() bool {
	return pf != nil && *pf >= 0
}

func main() {
	var p *MyFloat = nil
	ok := p.IsPositive()
	fmt.Println(ok)
}
