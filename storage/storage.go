package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mmbunde/todo/models"
)

func SetFilePath(fileName string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filePath := filepath.Join(homeDir, ".config", "todo", fileName)
	return filePath
}

func LoadTasks(fileName string) ([]models.Task, int, error) {
	var taskList []models.Task
	filePath := SetFilePath(fileName)
	taskData, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return taskList, 1, err
	} else if err != nil {
		fmt.Println("Unexpected error, quitting")
		os.Exit(1)
	}
	err = json.Unmarshal(taskData, &taskList)
	if err != nil {
		fmt.Println("Error parsing JSON, quitting")
		os.Exit(1)
	}
	if len(taskList) == 0 {
		return taskList, 1, nil
	}
	return taskList, taskList[len(taskList)-1].ID + 1, nil
}

func SaveTasks(fileName string, taskList []models.Task) {
	filePath := SetFilePath(fileName)
	taskData, err := json.Marshal(taskList)
	if err != nil {
		fmt.Println("Error encoding file to JSON, quitting")
		os.Exit(1)
	}
	err = os.WriteFile(filePath, taskData, 0644)
	if err != nil {
		fmt.Println("Error writing to file, quitting")
		os.Exit(1)
	}
}
