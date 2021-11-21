package cmd

import (
	"fmt"
	"strings"

	"github.com/mvenkatesh431/todo/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To Add a task to the Todo List",

	Run: func(cmd *cobra.Command, args []string) {
		todo := strings.Join(args, " ")
		_, err := db.CreateTodo(todo)
		if err != nil {
			fmt.Println("Failed to store Todo ", err)
		} else {
			fmt.Printf("Todo '%s' Added successfully to TodoList", todo)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
