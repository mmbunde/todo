package main

import (
	"bufio"
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
	message := "Would you like to add, list, delete, or complete a task or quit?"
loop:
	for {
		action := readInput(message)
		switch action {
		case "add":
			taskTitle = readInput("Name of the task: ")
			index := findTaskByTitle(tasks, taskTitle)
			if index == -1 || tasks[index].Done == true {
				tasks = append(tasks, Task{
					ID:    id,
					Title: taskTitle,
					Done:  false,
				})
				id++
			} else {
				fmt.Println("Task already exists and isn't complete")
			}
		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks are being tracked")
			} else {
				fmt.Printf("%-5s %-20s %s\n", "ID", "Title", "Done")
				for _, task := range tasks {
					check := "\u274c" //This is a red x
					if task.Done == true {
						check = "\u2705" //This is a green checkmark
					}
					fmt.Printf("%-5d %-20s %s\n", task.ID, task.Title, check)
				}
			}
		case "complete":
			taskTitle = readInput("Name of task completed: ")
			index := findTaskByTitle(tasks, taskTitle)
			if index != -1 {
				tasks[index].Done = true
			} else {
				fmt.Println("Task not found!")
			}
		case "delete":
			taskTitle = readInput("Name of the task to delete: ")
			index := findTaskByTitle(tasks, taskTitle)
			if index != -1 {
				tasks = slices.Delete(tasks, index, index+1)
			} else {
				fmt.Println("Task not found!")
			}
		case "quit", "exit":
			break loop
		default:
			fmt.Println("Please select add, list, delete, or quit")
		}
	}
}

// Helper functions to simplify the code
func readInput(prompt string) string {
	fmt.Println(prompt)
	userInput, _ := reader.ReadString('\n')
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
