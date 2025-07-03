package weather

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-agent/cmd/cliflags"
	"go-agent/internal/application"
	"go-agent/internal/domain"
	weatherService "go-agent/internal/infrastructure/weather"
)

var example = `# sfeir_cli weather Paris`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "weather",
		Short:   "Get weather for a city",
		Example: example,
		RunE: func(cmd *cobra.Command, args []string) error {

			flags, err := parseFlags(cmd)
			if err != nil {
				return err
			}

			weatherS := application.InitWeather(weatherService.NewWeatherClient())
			weather, err := weatherS.GetWeather(flags)
			if err != nil {
				return err
			}
			fmt.Println(weather)
			return nil
		},
	}

	flags := cmd.Flags()
	flags.String(cliflags.City, "", "The city to get the weather for")
	cmd.MarkFlagRequired(cliflags.City)
	return cmd
}

func parseFlags(cmd *cobra.Command) (domain.CheckWeatherOptions, error) {
	flags := cmd.Flags()

	city, err := flags.GetString(cliflags.City)
	if err != nil {
		return domain.CheckWeatherOptions{}, fmt.Errorf("failed to get %s flag: %w", cliflags.City, err)
	}

	return domain.CheckWeatherOptions{City: city}, nil
}
