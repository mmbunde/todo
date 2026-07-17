package cli

import (
	"fmt"
	"os"

	"github.com/mmbunde/todo/models"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [Name of the task]",
	Short: "Mark a task complete",

	Run: func(cmd *cobra.Command, args []string) {
		checkArgs(cmd, args)
		err := models.CompleteTask(db, args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
