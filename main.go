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
	message := "Would you like to add, list, delete, or quit a task?"
loop:
	for {
		action := readInput(message)
		switch action {
		case "add":
			taskTitle = readInput("Name of the task: ")
			tasks = append(tasks, Task{
				ID:    id,
				Title: taskTitle,
				Done:  true,
			})
			id++
		case "list":
			fmt.Printf("%-5s %-20s %s\n", "ID", "Title", "Done")
			for _, task := range tasks {
				check := "\u274c" //This is a red x
				if task.Done == true {
					check = "\u2705" //This is a green checkmark
				}
				fmt.Printf("%-5d %-20s %s\n", task.ID, task.Title, check)
			}
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
