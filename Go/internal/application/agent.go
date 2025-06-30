package application

import (
	"bufio"
	"context"
	"fmt"
	"github.com/anthropics/anthropic-sdk-go"
	"os"
)

func InitAgent(port LLMPort) AgentPort {
	return &Agent{
		llmPort: port,
	}
}

type Agent struct {
	llmPort LLMPort
}

func (a *Agent) Run(ctx context.Context) error {
	var conversation []anthropic.MessageParam

	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

	readUserInput := true
	for {
		if readUserInput {
			fmt.Print("\u001b[94mYou\u001b[0m: ")
			userInput, ok := a.getUserMessage()
			if !ok {
				break
			}
			userMessage := anthropic.NewUserMessage(anthropic.NewTextBlock(userInput))
			conversation = append(conversation, userMessage)
		}
		message, err := a.llmPort.RunInference(ctx, conversation)
		if err != nil {
			return err
		}
		conversation = append(conversation, message.ToParam())

		var toolResults []anthropic.ContentBlockParamUnion
		for _, content := range message.Content {
			switch content.Type {
			case "text":
				fmt.Printf("\u001b[93mClaude\u001b[0m: %s\n", content.Text)
			case "tool_use":
				result := a.llmPort.ExecuteTool(content.ID, content.Name, content.Input)
				toolResults = append(toolResults, result)
			}
		}
		if len(toolResults) == 0 {
			readUserInput = true
			continue
		}
		readUserInput = false
		conversation = append(conversation, anthropic.NewUserMessage(toolResults...))

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
