package cmd

import (
	"fmt"
	"strconv"

	"github.com/mvenkatesh431/todo/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To Delete a task from the Todo List ",
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			id, _ := strconv.Atoi(v)
			todos, err := db.ListTodos()
			if err != nil {
				fmt.Println("Failed to get the TodoList ", err)
			}
			if len(todos) == 0 {
				fmt.Println("Your TodoList is Empty!. Add few todos :)")
				return
			} else if (id <= 0) && (id > len(todos)) {
				fmt.Println("Invalid Todo ID")
				continue
			}

			// todos is zero indexed but user enters Ids starts with 1. So need to subtract 1
			todo := todos[id-1]

			// Now we got the 'todo' to delete from the DB
			err = db.DeleteTodo(todo.Key)
			if err != nil {
				fmt.Printf("Failed to delete %d Todo; error : %s", id, err)
			} else {
				fmt.Printf("Todo '%s' marked as Completed", todo.Value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
