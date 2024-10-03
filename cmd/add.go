package cmd

import (
	"github.com/antmordel/getting-started-cli/pkg"
	"github.com/spf13/cobra"
)

var todoText string

func init() {
	var addTodoCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds a new TODO item",
		Run:   addTodo,
	}

	addTodoCmd.Flags().StringVar(&todoText, "todo", "", "Text to add for the ToDo)")

	rootCmd.AddCommand(addTodoCmd)
}

func addTodo(cmd *cobra.Command, args []string) {
	if todoText == "" {
		todoText = pkg.AskForInputBlocking("Enter the chain canonical name: ")
	}

	db.New(todoText)
}
