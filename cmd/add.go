package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [Name of the task]",
	Short: "Adds a task",

	Run: func(cmd *cobra.Command, args []string) {
		checkArgs(cmd, args)
		// if len(args) == 0 {
		// 	cmd.Help()
		// 	os.Exit(1)
		// }
		_, err := db.Exec("INSERT INTO tasks (task_title) VALUES (?)", args[0])
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				fmt.Println("Task already exists")
			} else {
				fmt.Println(err)
			}
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
