package cmd

import (
	"go-agent/cmd/check_dns"
	"go-agent/cmd/check_host"
	"go-agent/cmd/check_port"
	"go-agent/cmd/llm_agent"

	"github.com/spf13/cobra"
)

func addCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(check_dns.New())
	rootCmd.AddCommand(check_host.New())
	rootCmd.AddCommand(check_port.New())
	rootCmd.AddCommand(llm_agent.New())
}
