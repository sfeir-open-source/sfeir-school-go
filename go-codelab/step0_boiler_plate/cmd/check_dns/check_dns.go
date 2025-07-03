package check_dns

import (
	"go-agent/cmd/cliflags"
	"go-agent/internal/domain"
	"github.com/spf13/cobra"
)

var example = `# sfeir_cli check-dns google.com`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check-dns",
		Short:   "Resolve DNS for a host",
		Example: example,
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	flags := cmd.Flags()
	flags.String(cliflags.Host, "", "Focus the host")
	cmd.MarkFlagRequired(cliflags.Host)
	return cmd
}

func parseFlags(cmd *cobra.Command) (domain.CheckHostOptions, error) {

	return domain.CheckHostOptions{}, nil
}
