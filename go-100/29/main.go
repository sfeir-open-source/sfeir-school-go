package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Le double de %v est %v\n", v, v*2)
	case string:
		fmt.Printf("%q a une longueur de %v octets\n", v, len(v))
	default:
		fmt.Printf("Je ne connais pas le type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
