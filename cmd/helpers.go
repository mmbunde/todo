package cmd

import (
	"database/sql"
	"os"

	"github.com/spf13/cobra"
)

func checkRowsAffected(taskResults sql.Result) (bool, error) {
	rows, err := taskResults.RowsAffected()
	if err != nil {
		return false, err
	}
	if rows == 0 {
		return false, nil
	}
	return true, nil
}

func checkArgs(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}
}
