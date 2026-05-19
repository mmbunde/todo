package cmd

import (
	"fmt"
	"os"

	"github.com/mmbunde/todo/models"
	"github.com/mmbunde/todo/storage"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all current tasks and their completion status.",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, _, err := storage.LoadTasks(taskFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		models.ListTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
