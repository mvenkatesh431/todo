package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To Add a task to the Todo List",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		for i, task := range args {
			fmt.Println(i, task)
		}
		fmt.Println(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
