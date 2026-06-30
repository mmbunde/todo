package cmd

import (
	"fmt"
	"os"

	"github.com/mmbunde/todo/models"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all current tasks and their completion status.",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := models.ListTask(db)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks being tracked")
			os.Exit(0)
		}
		fmt.Printf("%-5s %-20s %s\n", "ID", "Title", "Done")
		for _, task := range tasks {
			check := "\u274c" //This is a red x
			if task.Complete == 1 {
				check = "\u2705" //This is a green checkmark
			}
			fmt.Printf("%-5d %-20s %s\n", task.ID, task.TaskTitle, check)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
