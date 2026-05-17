package cmd

import (
	"github.com/mmbunde/todo/models"
	"github.com/mmbunde/todo/storage"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, id := storage.LoadTasks(taskFile)
		tasks, _ = models.AddTask(tasks, args[0], id)
		storage.SaveTasks(taskFile, tasks)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
