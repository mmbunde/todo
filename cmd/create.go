/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mmbunde/todo/models"
	"github.com/mmbunde/todo/storage"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the json file for your task list",

	Run: func(cmd *cobra.Command, args []string) {
		_, _, err := storage.LoadTasks(taskFile)
		if err != nil {
			storage.SaveTasks(taskFile, []models.Task{})
			filePath := storage.SetFilePath(taskFile)
			fmt.Println("Your file is saved at " + filePath)
		} else {
			fmt.Println("File already exists!")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
