package check_host

import (
	"go-agent/cmd/cliflags"
	"go-agent/internal/application"
	"go-agent/internal/domain"
	"fmt"
	"github.com/spf13/cobra"
)

var example = `# sfeir_cli check_host google.com`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check-host",
		Short:   "Ping a host",
		Example: example,
		RunE: func(cmd *cobra.Command, args []string) error {

			flags, err := parseFlags(cmd)
			if err != nil {
				return err
			}

			check := application.InitCheck()

			if err := check.PingHost(flags); err != nil {
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

func parseFlags(cmd *cobra.Command) (domain.CheckPingOptions, error) {
	flags := cmd.Flags()

	host, err := flags.GetString(cliflags.Host)
	if err != nil {
		return domain.CheckPingOptions{}, fmt.Errorf("failed to get %s flag: %w", cliflags.Host, err)
	}

	return domain.CheckPingOptions{Host: host}, nil
}
