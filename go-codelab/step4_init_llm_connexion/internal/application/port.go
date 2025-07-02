package application

import (
	"go-agent/internal/domain"
	"github.com/anthropics/anthropic-sdk-go"

	"context"
)

type AgentPort interface {
	Run(context.Context) error
}

type CheckerPort interface {
	ResolveDNS(opts domain.CheckHostOptions) error
}

// INFRA

type LLMPort interface {
	RunInference(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error)
}
