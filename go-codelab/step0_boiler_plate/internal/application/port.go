package application

import (
	"go-agent/internal/domain"
	"context"
)

type AgentPort interface {
	Run(context.Context) error
}

type CheckerPort interface {
	ResolveDNS(opts domain.CheckHostOptions) error
}
