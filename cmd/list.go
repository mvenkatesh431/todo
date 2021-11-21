package cmd

import (
	"fmt"

	"github.com/mvenkatesh431/todo/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "To List the tasks in the Todo List ",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := db.ListTodos()
		if err != nil {
			fmt.Println("Failed to get the TodoList ", err)
		}
		if len(todos) == 0 {
			fmt.Println("Your TodoList is Empty!. Add few todos :)")
			return
		}
		for i, todo := range todos {
			fmt.Printf("%d. %s\n", i+1, todo.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
