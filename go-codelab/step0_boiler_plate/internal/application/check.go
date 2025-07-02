package application

import (
	"go-agent/internal/domain"
)

func InitCheck() CheckerPort {
	return &Checker{}

}

type Checker struct{}

func (c Checker) ResolveDNS(opts domain.CheckHostOptions) error {

	return nil
}
