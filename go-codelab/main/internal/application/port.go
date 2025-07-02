package application

import (
	"go-agent/internal/domain"
	"context"
	"encoding/json"
	"github.com/anthropics/anthropic-sdk-go"
)

type AgentPort interface {
	Run(context.Context) error
}

type CheckerPort interface {
	ResolveDNS(opts domain.CheckHostOptions) error
	CheckPort(opts domain.CheckPortOptions) error
	PingHost(opts domain.CheckPingOptions) error
}

// INFRA

type LLMPort interface {
	RunInference(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error)
	ExecuteTool(id, name string, input json.RawMessage) anthropic.ContentBlockParamUnion
}
