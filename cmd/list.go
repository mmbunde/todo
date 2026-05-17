package cmd

import (
	"github.com/mmbunde/todo/models"
	"github.com/mmbunde/todo/storage"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all current tasks and their completion status.",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := storage.LoadTasks(taskFile)
		models.ListTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
