package cmd

import (
	"Go/cmd/check_dns"
	"Go/cmd/check_host"
	"Go/cmd/check_port"
	"Go/cmd/llm_agent"

	"github.com/spf13/cobra"
)

func addCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(check_dns.New())
	rootCmd.AddCommand(check_host.New())
	rootCmd.AddCommand(check_port.New())
	rootCmd.AddCommand(llm_agent.New())
}
