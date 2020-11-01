package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Logger définit une interface qui offre une unique méthode Log.
type Logger interface {
	Log(s string)
}

//FmtLogger est un logger simple qui écrit sur la sortie standard
type FmtLogger struct{}

//Log écritle message sur la sortie standard.
//Grâce à cette méthode, FmtLogger implémente l'interface Logger.
func (l *FmtLogger) Log(s string) {
	fmt.Println(s)
}

func recoverPanic(logger Logger) {
	r := recover()
	if r != nil {
		//Utilise le logger pour fournir une information sur la panique rattrapée.
		switch t := r.(type) {
		case string:
			logger.Log("PANIC: " + t)
		case error:
			logger.Log("PANIC: error is " + t.Error())
		default:
			logger.Log("PANIC: unknown error")
		}
	}
}

func main() {
	// Instance de notre logger
	logger := &FmtLogger{}
	// recoverPanic est appelée via defer de façon à ce qu'elle puisse rattraper la panique.
	defer recoverPanic(logger)

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
