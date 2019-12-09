package main

import (
	"errors"
	"math/rand"
	"time"
)

type Logger interface {
	Log(s string)
}

//TODO: créer un type FmtLogger qui implémente Logger

func recoverPanic(logger Logger) {
	r := recover() // permet de rattraper une panique. r est de type interface{}
	if r != nil {
		//TODO: utiliser le logger pour écrire ce que contient la panique (r)
		//Vous devrez gérer le cas où r est un string, où r est une error, et un cas par défaut où on loggera "PANIC: unknown error".
	}
}

func main() {
	//TODO instancier le FmtLogger
	//TODO appeler la fonction recoverPanic (de façon à ce qu'elle soit appelée même en cas de panique) en passant le logger en paramètre

	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(3)
	switch random {
	case 0:
		panic("Oups!")
	case 1:
		panic(errors.New("Zut"))
	default:
		panic(42)
	}
}
