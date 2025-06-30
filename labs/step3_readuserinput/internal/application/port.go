package application

import (
	"Go/internal/domain"
	"context"
)

type AgentPort interface {
	Run(context.Context) error
}

type CheckerPort interface {
	ResolveDNS(opts domain.CheckHostOptions) error
}
