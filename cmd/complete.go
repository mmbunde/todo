package cmd

import (
	"database/sql"
	"errors"
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
		var complete int
		err := db.QueryRow("SELECT complete FROM tasks WHERE task_title = (?)", args[0]).Scan(&complete)
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Task doesn't exist")
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if complete == 1 {
			fmt.Println("Task is already complete")
			os.Exit(1)
		}
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
