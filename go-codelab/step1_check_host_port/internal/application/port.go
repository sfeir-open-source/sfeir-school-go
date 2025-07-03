package application

import (
	"go-agent/internal/domain"
)

type CheckerPort interface {
	ResolveDNS(opts domain.CheckHostOptions) error
}
