package anthropic

import (
	"Go/internal/application"
	"context"
	"github.com/anthropics/anthropic-sdk-go"
)

func NewLLM() application.LLMPort {
	client := anthropic.NewClient()
	return &adapter{client: &client}
}

type adapter struct {
	client *anthropic.Client
}

func (a *adapter) RunInference(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error) {

	message, err := a.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_7SonnetLatest,
		MaxTokens: int64(1024),
		Messages:  conversation,
	})
	return message, err
}
