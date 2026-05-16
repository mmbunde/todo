package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var usage = "Usage:\n" +
	"todo [JSON filepath] list\n" +
	"todo [JSON filepath] add|delete|complete [Task Title]\n"

func main() {
	var err error

	fileName, action, taskTitle := validateArgs()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filePath := filepath.Join(homeDir, ".config", "todo", fileName)
	tasks, id := loadTasks(filePath)
	switch action {
	case "add":
		tasks, id = addTask(tasks, taskTitle, id)
	case "list":
		listTasks(tasks)
	case "complete":
		tasks, err = completeTask(tasks, taskTitle)
		if err != nil {
			fmt.Println(err)
		}
	case "delete":
		tasks, err = deleteTask(tasks, taskTitle)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println(usage)
	}
	saveTasks(filePath, tasks)
}

func validateArgs() (string, string, string) {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}
	fileName := os.Args[1]
	action := os.Args[2]
	if !strings.EqualFold(filepath.Ext(fileName), ".json") {
		fmt.Println(usage)
		os.Exit(1)
	} else if len(os.Args) == 3 && os.Args[2] != "list" {
		fmt.Println(usage)
		os.Exit(1)
	} else if len(os.Args) == 4 && os.Args[2] != "add" && os.Args[2] != "delete" && os.Args[2] != "complete" {
		fmt.Println(usage)
		os.Exit(1)
	} else if len(os.Args) > 4 {
		fmt.Println(usage)
		os.Exit(1)
	}
	if len(os.Args) == 4 {
		taskTitle := os.Args[3]
		return fileName, action, taskTitle
	}
	return fileName, action, ""
}
