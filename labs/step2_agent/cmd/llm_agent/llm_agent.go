package llm_agent

import (
	"Go/internal/application"
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
			agent := application.InitAgent()
			err := agent.Run(context.TODO())
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}

		},
	}

	return cmd
}
