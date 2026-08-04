package cli

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mmbunde/todo/storage"
	"github.com/spf13/cobra"
)

var taskFile string
var db *sql.DB

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
todo -f tasks.db add "Learn Go"`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch cmd.Name() {
		case "completion", "__complete", "__completeNoDesc":
			return
		}
		if taskFile == "" {
			cmd.Help()
			os.Exit(1)
		}
		var err error
		db, err = storage.InitDB(taskFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if db != nil {
			err := db.Close()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&taskFile, "file", "f", "", "SQLite DB to store tasks")
	rootCmd.RegisterFlagCompletionFunc("file",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"db"}, cobra.ShellCompDirectiveFilterFileExt
		})
}
