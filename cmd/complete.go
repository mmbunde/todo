package cmd

import (
	"fmt"
	"os"

	"github.com/mmbunde/todo/models"
	"github.com/mmbunde/todo/storage"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task complete",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		tasks, _ := storage.LoadTasks(taskFile)
		tasks, err = models.CompleteTask(tasks, args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		storage.SaveTasks(taskFile, tasks)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
