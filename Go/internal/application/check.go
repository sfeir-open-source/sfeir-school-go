package application

import (
	"Go/internal/domain"
	"fmt"
	"net"
	"time"
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

func (c Checker) PingHost(opts domain.CheckPingOptions) error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", opts.Host, 80), 2*time.Second)
	if err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	defer conn.Close()
	fmt.Printf("TCP ping to %s succeeded\n", opts.Host)
	return nil
}

func (c Checker) CheckPort(opts domain.CheckPortOptions) error {
	conn, err := net.DialTimeout("tcp", opts.Host, 2*time.Second)
	if err != nil {
		return fmt.Errorf("port check failed: %w", err)
	}
	defer conn.Close()
	fmt.Printf("Port %s is open\n", opts.Host)
	return nil
}
