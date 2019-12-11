package main

import "fmt"

type Age int

type Person struct {
	name   string
	age    Age
	friend *Person
}

func main() {
	john := Person{"John", 25, nil}
	pJohn := &john
	pJohn.age++

	//Indiquez true ou false en troisième argument
	assert("Q11", john.age == 25, false)
	assert("Q12", john.age == 26, true)
	assert("Q13", pJohn.age == 25, false)
	assert("Q14", pJohn.age == 26, true)

	pAgeJ := &(john.age)
	bob := Person{"Bob", 26, &john}
	pAgeB := &(bob.age)

	//Indiquez true ou false en troisième argument
	assert("Q2", pAgeJ == pAgeB, false)

	//Indiquez true ou false en troisième argument
	assert("Q31", *pAgeJ == john.age, true)
	assert("Q32", *pAgeJ == bob.age, true)

	*pAgeJ = 10

	//Indiquez true ou false en troisième argument
	assert("Q41", *pAgeJ == john.age, true)
	assert("Q42", *pAgeJ == bob.age, false)

	john.age = 12

	//Indiquez true ou false en troisième argument
	assert("Q51", *pAgeJ == john.age, true)
	assert("Q52", *pAgeJ == bob.age, false)

	john.age = *pAgeB

	//Indiquez true ou false en troisième argument
	assert("Q61", *pAgeJ == john.age, true)
	assert("Q62", *pAgeJ == bob.age, true)

	*pAgeB = 18

	//Indiquez true ou false en troisième argument
	assert("Q71", john.age == 18, false)
	assert("Q72", bob.age == 18, true)

	bob.friend.age = *pAgeB + 1

	//Indiquez true ou false en troisième argument
	assert("Q81", john.age == 19, true)
	assert("Q82", bob.age == 19, false)

	bob.friend = &bob
	bob.friend.age = 20

	//Indiquez true ou false en troisième argument
	assert("Q91", john.age == 20, false)
	assert("Q92", bob.age == 20, true)

	bob.friend = &bob
	pFriend := bob.friend
	bob.friend.friend = &john

	//Indiquez true ou false en troisième argument
	assert("Q101", bob.friend == pFriend, false)
	assert("Q102", bob.friend == pJohn, true)

	eric := john

	//Indiquez true ou false en troisième argument
	assert("Q111", eric == john, true)
	assert("Q112", &eric == &john, false)
	assert("Q113", *&eric == *&john, true)

	eric.name = "Eric"

	//Indiquez true ou false en troisième argument
	assert("Q121", john.name == "Eric", false)
	assert("Q122", eric == john, false)
	assert("Q123", &eric == &john, false)
	assert("Q124", *&eric == *&john, false)

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
