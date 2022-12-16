package main

import "fmt"

const REPLACEME = "replaceme"

// Structure
type Age int

type Person struct {
	name   string
	age    Age
	friend *Person
}

func main() {

	john := Person{"John", 25, nil}
	pJohn := &john

	// En Go, l'indirection via le pointeur est transparente.
	// Cela signifie que cette ligne de code est équivalente à (*pJohn).age++
	pJohn.age++

	// Indiquez true ou false en troisième argument
	assert("Q11", john.age == 25, REPLACEME)
	assert("Q12", john.age == 26, REPLACEME)
	assert("Q13", pJohn.age == 25, REPLACEME)
	assert("Q14", pJohn.age == 26, REPLACEME)

	pAgeJ := &(john.age)
	bob := Person{"Bob", 26, &john}
	pAgeB := &(bob.age)

	// Indiquez true ou false en troisième argument
	assert("Q2", pAgeJ == pAgeB, REPLACEME)

	// Indiquez true ou false en troisième argument
	assert("Q31", *pAgeJ == john.age, REPLACEME)
	assert("Q32", *pAgeJ == bob.age, REPLACEME)

	*pAgeJ = 10

	// Indiquez true ou false en troisième argument
	assert("Q41", *pAgeJ == john.age, REPLACEME)
	assert("Q42", *pAgeJ == bob.age, REPLACEME)

	john.age = 12

	// Indiquez true ou false en troisième argument
	assert("Q51", *pAgeJ == john.age, REPLACEME)
	assert("Q52", *pAgeJ == bob.age, REPLACEME)

	john.age = *pAgeB

	// Indiquez true ou false en troisième argument
	assert("Q61", *pAgeJ == john.age, REPLACEME)
	assert("Q62", *pAgeJ == bob.age, REPLACEME)

	*pAgeB = 18

	//Indiquez true ou false en troisième argument
	assert("Q71", john.age == 18, REPLACEME)
	assert("Q72", bob.age == 18, REPLACEME)

	bob.friend.age = *pAgeB + 1

	// Indiquez true ou false en troisième argument
	assert("Q81", john.age == 19, REPLACEME)
	assert("Q82", bob.age == 19, REPLACEME)

	bob.friend = &bob
	bob.friend.age = 20

	// Indiquez true ou false en troisième argument
	assert("Q91", john.age == 20, REPLACEME)
	assert("Q92", bob.age == 20, REPLACEME)

	bob.friend = &bob
	pFriend := bob.friend
	bob.friend.friend = &john

	//Indiquez true ou false en troisième argument
	assert("Q101", bob.friend == pFriend, REPLACEME)
	assert("Q102", bob.friend == pJohn, REPLACEME)

	eric := john

	// Indiquez true ou false en troisième argument
	assert("Q111", eric == john, REPLACEME)
	assert("Q112", &eric == &john, REPLACEME)
	assert("Q113", *&eric == *&john, REPLACEME)

	eric.name = "Eric"

	// Indiquez true ou false en troisième argument
	assert("Q121", john.name == "Eric", REPLACEME)
	assert("Q122", eric == john, REPLACEME)
	assert("Q123", &eric == &john, REPLACEME)
	assert("Q124", *&eric == *&john, REPLACEME)

}

func assert(message string, value bool, expected interface{}) {
	ve, ok := expected.(bool)
	if !ok {
		fmt.Printf("%5s : %s\n", message, "Missing answer")
		return
	}

	if value != ve {
		fmt.Printf("%5s : %s\n", message, "Wrong answer")
	} else {
		fmt.Printf("%5s : %s\n", message, "Correct answer")
	}
}
