package cmd

import (
	"github.com/spf13/cobra"
)

var CLIName = "getting-started-cli"

var rootCmd = &cobra.Command{
	Use:   CLIName,
	Short: "Getting Started CLI",
	Long:  "A CLI tool for demonstrating how to create a CLI tool using Golang.",
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
