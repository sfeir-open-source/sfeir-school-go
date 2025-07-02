package cmd

import (
	"github.com/spf13/cobra/doc"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "dev" // Default version for local builds
)

var rootCmd = &cobra.Command{
	Use:     "sfeir_cli",
	Short:   "helper for operation",
	Long:    "helper for operational actions",
	Version: version,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

var genDocsCmd = &cobra.Command{
	Use:   "gen-docs",
	Short: "Generate markdown documentation for the CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "./docs/cmd"
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		return doc.GenMarkdownTree(rootCmd, dir)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	addCommands(rootCmd)
	rootCmd.AddCommand(genDocsCmd)
}

func initConfig() {

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}
}
