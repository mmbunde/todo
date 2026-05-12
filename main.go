package main

import (
	"bufio"
	"fmt"
	"os"
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
