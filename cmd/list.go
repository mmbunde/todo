package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all current tasks and their completion status.",

	Run: func(cmd *cobra.Command, args []string) {
		var id int
		var taskTitle string
		var complete int
		rows, err := db.Query("SELECT * FROM tasks")
		defer rows.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		found := false
		for rows.Next() {
			if !found {
				fmt.Printf("%-5s %-20s %s\n", "ID", "Title", "Done")
				found = true
			}
			rows.Scan(&id, &taskTitle, &complete)
			check := "\u274c"
			if complete == 1 {
				check = "\u2705"
			}
			fmt.Printf("%-5d %-20s %s\n", id, taskTitle, check)
		}
		if !found {
			fmt.Println("No tasks being tracked")
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
