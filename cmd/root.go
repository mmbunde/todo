package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var taskFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A CLI tool to manage your tasks",
	Long: `A CLI tool to manage your tasks

todo -f <file> create
todo -f <file> list
todo -f <file> add <task title>
todo -f <file> complete <task title>
todo -f <file> delete <task title>

Example:
todo -f tasks.json add "Learn Go"`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch cmd.Name() {
		case "completion", "__complete", "__completeNoDesc":
			return
		}
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
	rootCmd.PersistentFlags().StringVarP(&taskFile, "file", "f", "", "JSON file to store tasks")
	rootCmd.RegisterFlagCompletionFunc("file",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"json"}, cobra.ShellCompDirectiveFilterFileExt
		})
}
