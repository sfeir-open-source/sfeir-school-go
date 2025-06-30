package check_dns

import (
	"Go/cmd/cliflags"
	"Go/internal/application"
	"Go/internal/domain"
	"fmt"
	"github.com/spf13/cobra"
)

var example = `# sfeir_cli check-dns google.com`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check-dns",
		Short:   "Resolve DNS for a host",
		Example: example,
		RunE: func(cmd *cobra.Command, args []string) error {

			flags, err := parseFlags(cmd)
			if err != nil {
				return err
			}

			check := application.InitCheck()
			if err := check.ResolveDNS(flags); err != nil {
				fmt.Println("Error:", err)
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.String(cliflags.Host, "", "Focus the host")
	cmd.MarkFlagRequired(cliflags.Host)
	return cmd
}

func parseFlags(cmd *cobra.Command) (domain.CheckHostOptions, error) {
	flags := cmd.Flags()

	host, err := flags.GetString(cliflags.Host)
	if err != nil {
		return domain.CheckHostOptions{}, fmt.Errorf("failed to get %s flag: %w", cliflags.Host, err)
	}

	return domain.CheckHostOptions{Host: host}, nil
}
