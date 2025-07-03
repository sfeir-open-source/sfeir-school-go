package application

import (
	"context"
	"encoding/json"
	"github.com/anthropics/anthropic-sdk-go"
	"go-agent/internal/domain"
)

type AgentPort interface {
	Run(context.Context) error
}

type CheckerPort interface {
	ResolveDNS(opts domain.CheckHostOptions) error
	CheckPort(opts domain.CheckPortOptions) error
	PingHost(opts domain.CheckPingOptions) error
}

type WeatherPort interface {
	GetWeather(opts domain.CheckWeatherOptions) (string, error)
}

// INFRA

type LLMPort interface {
	RunInference(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error)
	ExecuteTool(id, name string, input json.RawMessage) anthropic.ContentBlockParamUnion
}

type WeatherServicePort interface {
	GetWeather(opts domain.CheckWeatherOptions) (string, error)
}
