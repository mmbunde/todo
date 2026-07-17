package cli

import (
	"fmt"
	"os"

	"github.com/mmbunde/todo/models"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [Name of the task]",
	Short: "Delete a task",

	Run: func(cmd *cobra.Command, args []string) {
		checkArgs(cmd, args)
		err := models.DeleteTask(db, args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
