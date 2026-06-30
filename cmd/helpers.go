package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func checkArgs(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}
}
