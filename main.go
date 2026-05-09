package main

import "fmt"

func main() {
	id := 1
	var tasks []Task
	var input string
	message := "Would you like to add, list, or delete a task?"
	fmt.Println(message)
	for {
		fmt.Scanln(&input)
		if input == "yes" {
			fmt.Scanln(&input)
			tasks = append(tasks, Task{
				ID:    id,
				Title: input,
				Done:  false,
			})
			fmt.Println(message)
			id++
		}
		if input == "no" {
			fmt.Println(message)
		}
		if input == "list" {
			fmt.Println("Tasks: ", tasks)
		}

		if input == "quit" || input == "exit" {
			break
		}
	}
}
