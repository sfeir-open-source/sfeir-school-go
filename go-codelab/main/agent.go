package main

import (
	"go-agent/cmd"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	cmd.Execute()
}
