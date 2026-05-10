package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	id := 1
	var tasks []Task
	var taskTitle string
	reader := bufio.NewReader(os.Stdin)
	message := "Would you like to add, list, or delete a task?"
loop:
	for {
		fmt.Println(message)
		taskTitle, _ = reader.ReadString('\n')
		switch strings.TrimSpace(taskTitle) {
		case "add":
			fmt.Print("Name of the task: ")
			taskTitle, _ = reader.ReadString('\n')
			tasks = append(tasks, Task{
				ID:    id,
				Title: strings.TrimSpace(taskTitle),
				Done:  false,
			})
			id++
		case "list":
			fmt.Println("Tasks: ", tasks)
		case "delete":
			fmt.Print("Name of the task to delete: ")
			taskTitle, _ = reader.ReadString('\n')
			for i, task := range tasks {
				if task.Title == strings.TrimSpace(taskTitle) {
					tasks = slices.Delete(tasks, i, i+1)
				}
			}
		case "quit", "exit":
			break loop
		default:
			fmt.Println("Unknown error")
		}
	}
}
