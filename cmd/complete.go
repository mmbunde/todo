package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [Name of the task]",
	Short: "Mark a task complete",

	Run: func(cmd *cobra.Command, args []string) {
		checkArgs(cmd, args)
		// if len(args) == 0 {
		// 	cmd.Help()
		// 	os.Exit(1)
		// }
		taskResults, err := db.Exec("UPDATE tasks SET complete = 1 WHERE task_title = (?)", args[0])
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
			fmt.Println("No task found to complete")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
