package main

import (
	"fmt"
	"slices"
)

func main() {
	id := 1
	var tasks []Task
	var input string
	message := "Would you like to add, list, or delete a task?"
loop:
	for {
		fmt.Println(message)
		fmt.Scanln(&input)
		switch input {
		case "add":
			fmt.Print("Name of the task: ")
			fmt.Scanln(&input)
			tasks = append(tasks, Task{
				ID:    id,
				Title: input,
				Done:  false,
			})
			id++
		case "list":
			fmt.Println("Tasks: ", tasks)
		case "delete":
			fmt.Print("Name of the task to be detele: ")
			fmt.Scanln(&input)
			for i, task := range tasks {
				if task.Title == input {
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
