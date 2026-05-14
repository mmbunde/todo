package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func loadTasks(filePath string) ([]Task, int) {
	var taskList []Task
	taskData, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		fmt.Println("File will be created when you quit")
		return taskList, 1
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
		return taskList, 1
	}
	return taskList, taskList[len(taskList)-1].ID + 1
}

func saveTasks(filePath string, taskList []Task) {
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

func addTask(taskList []Task, taskTitle string, id int) ([]Task, int) {
	index := findTaskByTitle(taskList, taskTitle)
	if index == -1 || taskList[index].Done == true {
		taskList = append(taskList, Task{
			ID:    id,
			Title: taskTitle,
			Done:  false,
		})
		id++
	} else {
		fmt.Println("Task already exists and isn't complete")
	}
	return taskList, id
}

func completeTask(taskList []Task, taskTitle string) ([]Task, error) {
	index := findTaskByTitle(taskList, taskTitle)
	if index != -1 {
		taskList[index].Done = true
		return taskList, nil
	}
	return taskList, errors.New("Task not found!")
}

func deleteTask(taskList []Task, taskTitle string) ([]Task, error) {
	index := findTaskByTitle(taskList, taskTitle)
	if index != -1 {
		return slices.Delete(taskList, index, index+1), nil
	}
	return taskList, errors.New("Task not found!")
}

func listTasks(taskList []Task) {
	if len(taskList) == 0 {
		fmt.Println("No tasks are being tracked")
	} else {
		fmt.Printf("%-5s %-20s %s\n", "ID", "Title", "Done")
		for _, task := range taskList {
			check := "\u274c" //This is a red x
			if task.Done == true {
				check = "\u2705" //This is a green checkmark
			}
			fmt.Printf("%-5d %-20s %s\n", task.ID, task.Title, check)
		}
	}
}
