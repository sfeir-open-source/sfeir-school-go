package anthropic

import (
	"Go/internal/application"
	"context"
	"encoding/json"
	"fmt"
	"github.com/anthropics/anthropic-sdk-go"
)

func NewLLM(tools []ToolDefinition) application.LLMPort {
	client := anthropic.NewClient()
	var anthropicTools []anthropic.ToolUnionParam

	for _, tool := range tools {
		anthropicTools = append(anthropicTools, anthropic.ToolUnionParam{
			OfTool: &anthropic.ToolParam{
				Name:        tool.Name,
				Description: anthropic.String(tool.Description),
				InputSchema: tool.InputSchema,
			},
		})
	}
	return &adapter{client: &client, tools: anthropicTools, rawTools: tools}
}

type adapter struct {
	client   *anthropic.Client
	tools    []anthropic.ToolUnionParam
	rawTools []ToolDefinition
}

func (a *adapter) RunInference(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error) {

	message, err := a.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_7SonnetLatest,
		MaxTokens: int64(1024),
		Messages:  conversation,
		Tools:     a.tools,
	})
	return message, err
}

func (a *adapter) ExecuteTool(id, name string, input json.RawMessage) anthropic.ContentBlockParamUnion {
	var toolDef ToolDefinition
	var found bool
	for _, tool := range a.rawTools {
		if tool.Name == name {
			toolDef = tool
			found = true
			break
		}
	}
	if !found {
		return anthropic.NewToolResultBlock(id, "tool not found", true)
	}

	fmt.Printf("\u001b[92mtool\u001b[0m: %s(%s)\n", name, input)
	response, err := toolDef.Function(input)
	if err != nil {
		return anthropic.NewToolResultBlock(id, err.Error(), true)
	}
	return anthropic.NewToolResultBlock(id, response, false)
}
