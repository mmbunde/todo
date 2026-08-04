package cli

import (
	"fmt"
	"os"

	"github.com/mmbunde/todo/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [Name of the task]",
	Short: "Adds a task",

	Run: func(cmd *cobra.Command, args []string) {
		checkArgs(cmd, args)
		err := models.AddTask(db, args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
