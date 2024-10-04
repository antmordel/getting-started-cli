package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var listTodoCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all the TODOs",
		Run:   listTodoCmd,
	}

	rootCmd.AddCommand(listTodoCmd)
}

func listTodoCmd(cmd *cobra.Command, args []string) {

	todos := db.All()

	for _, todo := range todos {
		if todo.Done {
			fmt.Printf("[x] %s\n", todo.Text)
		} else {
			fmt.Printf("[ ] %s\n", todo.Text)
		}
	}
}
