package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
}

func (p *person) fullname() string {
	if p == nil {
		return "Anonyme"
	}
	return p.Firstname + " " + p.Lastname
}

func main() {
	jack := &person{"Jack", "Sparrow"}
	fmt.Println(jack.fullname())
	jack = nil
	fmt.Println(jack.fullname())
}
