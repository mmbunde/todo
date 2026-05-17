package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var taskFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A tool to help you keep track of tasks",
	Long: `A tool to help you keep track of tasks

	Usage: 
	todo --file tasks.json list
	todo --file tasks.json add|delete|complete [Task Title]`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if taskFile == "" {
			cmd.Help()
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&taskFile, "file", "", "JSON file to store tasks ($HOME/.config/todo/tasks.json)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
