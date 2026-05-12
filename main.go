package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	id := 1
	var tasks []Task
	var taskTitle string
	var err error
	message := "Would you like to add, list, delete, or complete a task or quit?"
loop:
	for {
		action := readInput(message)
		switch action {
		case "add":
			taskTitle = readInput("Name of the task: ")
			tasks, id = addTask(tasks, taskTitle, id)
		case "list":
			listTasks(tasks)
		case "complete":
			taskTitle = readInput("Name of task completed: ")
			tasks, err = completeTask(tasks, taskTitle)
			if err != nil {
				fmt.Println(err)
			}
		case "delete":
			taskTitle = readInput("Name of the task to delete: ")
			tasks, err = deleteTask(tasks, taskTitle)
			if err != nil {
				fmt.Println(err)
			}
		case "quit", "exit":
			break loop
		default:
			fmt.Println("Please select add, list, delete, or quit")
		}
	}
}

// Helper functions
func readInput(prompt string) string {
	fmt.Println(prompt)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Unexpected Error, quitting")
		os.Exit(1)
	}
	return strings.TrimSpace(userInput)
}

func findTaskByTitle(taskList []Task, taskTitle string) int {
	for i, task := range taskList {
		if strings.EqualFold(task.Title, taskTitle) {
			return i
		}
	}
	return -1
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
