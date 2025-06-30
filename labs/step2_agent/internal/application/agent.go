package application

import (
	"context"
	"fmt"
)

func InitAgent() AgentPort {
	return &Agent{}
}

type Agent struct {
}

func (a *Agent) Run(ctx context.Context) error {

	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

	for {
		fmt.Printf("\u001b[93mClaude\u001b[0m: %s\n")
	}

	return nil
}
