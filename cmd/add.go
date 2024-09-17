package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var addTodoCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds a new TODO item",
		Run:   addTodo,
	}

	rootCmd.AddCommand(addTodoCmd)
}

func addTodo(cmd *cobra.Command, args []string) {
	fmt.Println("Adding a new TODO item")
}
