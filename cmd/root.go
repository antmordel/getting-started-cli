package cmd

import (
	"github.com/antmordel/getting-started-cli/internal"
	"github.com/spf13/cobra"
)

var CLIName = "todocli"

var db = internal.NewInMemoryTodoDB()

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
