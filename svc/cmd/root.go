package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "interview",
		Short: "interview",
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}

//nolint:gochecknoinits // framework feature
func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file (default are config.yaml and secrets.yaml)")
}
