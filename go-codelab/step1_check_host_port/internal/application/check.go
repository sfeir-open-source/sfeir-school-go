package application

import (
	"go-agent/internal/domain"
	"fmt"
	"net"
)

func InitCheck() CheckerPort {
	return &Checker{}

}

type Checker struct{}

func (c Checker) ResolveDNS(opts domain.CheckHostOptions) error {
	ips, err := net.LookupHost(opts.Host)
	if err != nil {
		return fmt.Errorf("DNS lookup failed: %w", err)
	}
	fmt.Printf("DNS resolution for %s:\n", opts.Host)
	for _, ip := range ips {
		fmt.Println(" -", ip)
	}
	return nil
}
