package cmd

import (
	"fmt"
	"os"

	"github.com/mmbunde/todo/models"
	"github.com/mmbunde/todo/storage"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		tasks, _ := storage.LoadTasks(taskFile)
		tasks, err = models.DeleteTask(tasks, args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		storage.SaveTasks(taskFile, tasks)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
