package application

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func InitAgent() AgentPort {
	return &Agent{}
}

type Agent struct {
}

func (a *Agent) Run(ctx context.Context) error {

	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

	for {
		fmt.Print("\u001b[94mYou\u001b[0m: ")
		userInput, ok := a.getUserMessage()
		if !ok {
			break
		}
		fmt.Printf("\u001b[93mClaude\u001b[0m: %s\n", userInput)
	}

	return nil
}

func (a *Agent) getUserMessage() (string, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", false
	}
	return scanner.Text(), true
}
