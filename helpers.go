package main

import (
	"fmt"
	"os"
	"strings"
)

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
