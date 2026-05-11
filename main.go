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
	message := "Would you like to add, list, or delete a task?"
loop:
	for {
		action := readInput(message)
		switch action {
		case "add":
			taskTitle = readInput("Name of the task: ")
			tasks = append(tasks, Task{
				ID:    id,
				Title: taskTitle,
				Done:  false,
			})
			id++
		case "list":
			fmt.Println("Tasks: ", tasks)
		case "delete":
			taskTitle = readInput("Name of the task to delete: ")
			for i, task := range tasks {
				if task.Title == taskTitle {
					tasks = slices.Delete(tasks, i, i+1)
				}
			}
		case "quit", "exit":
			break loop
		default:
			fmt.Println("Please select add, list or, delete")
		}
	}
}

// Helper functions to simplify the code
func readInput(prompt string) string {
	fmt.Println(prompt)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}
