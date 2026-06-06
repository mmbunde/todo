package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [Name of the task]",
	Short: "Delete a task",

	Run: func(cmd *cobra.Command, args []string) {
		checkArgs(cmd, args)
		// if len(args) == 0 {
		// 	cmd.Help()
		// 	os.Exit(1)
		// }
		taskResults, err := db.Exec("DELETE FROM tasks WHERE task_title = (?)", args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		check, err := checkRowsAffected(taskResults)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if !check {
			fmt.Println("No task found to delete")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
