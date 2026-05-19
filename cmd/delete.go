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
	Use:   "delete [Name of the task]",
	Short: "Delete a task",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		tasks, _, err := storage.LoadTasks(taskFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
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
