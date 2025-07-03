package cmd

import (
	"go-agent/cmd/check_dns"
	"github.com/spf13/cobra"
)

func addCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(check_dns.New())
}
