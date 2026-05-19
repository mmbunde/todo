package cmd

import (
	"fmt"
	"os"

	"github.com/mmbunde/todo/models"
	"github.com/mmbunde/todo/storage"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [Name of the task]",
	Short: "Adds a task",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, id, err := storage.LoadTasks(taskFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		tasks, _ = models.AddTask(tasks, args[0], id)
		storage.SaveTasks(taskFile, tasks)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
