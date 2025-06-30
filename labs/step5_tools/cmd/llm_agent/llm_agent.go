package llm_agent

import (
	"Go/internal/application"
	anthropicLocal "Go/internal/infrastructure/anthropic"
	"Go/internal/infrastructure/tools"
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

var example = `# sfeir_cli llm-agent`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "llm-agent",
		Short:   "Run the LLM agent",
		Example: example,
		Run: func(cmd *cobra.Command, args []string) {

			definitions := []anthropicLocal.ToolDefinition{
				tools.GetWeatherDefinition,
			}
			agent := application.InitAgent(anthropicLocal.NewLLM(definitions))

			err := agent.Run(context.TODO())
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}

		},
	}

	return cmd
}
