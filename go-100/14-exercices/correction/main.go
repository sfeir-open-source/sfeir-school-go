package main

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
	assert("Q1", john.age == 25, false)
	assert("Q1", john.age == 26, true)
	assert("Q1", pJohn.age == 25, false)
	assert("Q1", pJohn.age == 26, true)

	pAgeJ := &(john.age)
	bob := Person{"Bob", 26, &john}
	pAgeB := &(bob.age)

	//Indiquez true ou false en troisième argument
	assert("Q2", pAgeJ == pAgeB, false)

	//Indiquez true ou false en troisième argument
	assert("Q3", *pAgeJ == john.age, true)
	assert("Q3", *pAgeJ == bob.age, true)

	*pAgeJ = 10

	//Indiquez true ou false en troisième argument
	assert("Q4", *pAgeJ == john.age, true)
	assert("Q4", *pAgeJ == bob.age, false)

	john.age = 12

	//Indiquez true ou false en troisième argument
	assert("Q5", *pAgeJ == john.age, true)
	assert("Q5", *pAgeJ == bob.age, false)

	john.age = *pAgeB

	//Indiquez true ou false en troisième argument
	assert("Q6", *pAgeJ == john.age, true)
	assert("Q6", *pAgeJ == bob.age, true)

	*pAgeB = 18

	//Indiquez true ou false en troisième argument
	assert("Q7", john.age == 18, false)
	assert("Q7", bob.age == 18, true)

	bob.friend.age = *pAgeB + 1

	//Indiquez true ou false en troisième argument
	assert("Q8", john.age == 19, true)
	assert("Q8", bob.age == 19, false)

	bob.friend = &bob
	bob.friend.age = 20

	//Indiquez true ou false en troisième argument
	assert("Q8", john.age == 20, false)
	assert("Q8", bob.age == 20, true)

	bob.friend = &bob
	pFriend := bob.friend
	bob.friend.friend = &john

	//Indiquez true ou false en troisième argument
	assert("Q9", bob.friend == pFriend, false)
	assert("Q9", bob.friend == pJohn, true)

	eric := john

	//Indiquez true ou false en troisième argument
	assert("Q10", eric == john, true)
	assert("Q10", &eric == &john, false)
	assert("Q10", *&eric == *&john, true)

	eric.name = "Eric"

	//Indiquez true ou false en troisième argument
	assert("Q11", john.name == "Eric", false)
	assert("Q11", eric == john, false)
	assert("Q11", &eric == &john, false)
	assert("Q11", *&eric == *&john, false)

}

func assert(message string, value, expected bool) {
	if value != expected {
		panic(message)
	}
}
