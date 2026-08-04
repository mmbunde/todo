package models

import (
	"database/sql"
	"errors"
	"strings"
)

var (
	ErrTaskExist        = errors.New("task already exists")
	ErrTaskNotFound     = errors.New("task doesn't exist")
	ErrTaskComplete     = errors.New("task is already complete")
	ErrTaskUpdateFailed = errors.New("task could not be updated")
)

type Task struct {
	ID        int
	TaskTitle string
	Complete  int
}

func AddTask(taskDB *sql.DB, taskTitle string) error {
	_, err := taskDB.Exec("INSERT INTO tasks (task_title) VALUES (?)", taskTitle)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrTaskExist
		}
		return err
	}
	return nil
}

func ListTask(taskDB *sql.DB) ([]Task, error) {
	var tasks []Task
	rows, err := taskDB.Query("SELECT * FROM tasks")
	if err != nil {
		return tasks, err
	}
	defer rows.Close()
	for rows.Next() {
		var task Task // Need a single Task to hold the queried data
		rows.Scan(&task.ID, &task.TaskTitle, &task.Complete)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CompleteTask(taskDB *sql.DB, args string) error {
	var complete int
	err := taskDB.QueryRow("SELECT complete FROM tasks WHERE task_title = (?)", args).Scan(&complete)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrTaskNotFound
	}
	if err != nil {
		return err
	}
	if complete == 1 {
		err = errors.New("Task is already complete")
		return ErrTaskComplete
	}
	taskResults, err := taskDB.Exec("UPDATE tasks SET complete = 1 WHERE task_title = (?)", args)
	if err != nil {
		return err
	}
	check, err := checkRowsAffected(taskResults)
	if err != nil {
		return err
	}
	if !check {
		return ErrTaskUpdateFailed
	}
	return nil
}

func DeleteTask(taskDB *sql.DB, args string) error {
	taskResults, err := taskDB.Exec("DELETE FROM tasks WHERE task_title = (?)", args)
	if err != nil {
		return err
	}
	check, err := checkRowsAffected(taskResults)
	if err != nil {
		return err
	}
	if !check {
		return ErrTaskNotFound
	}
	return nil
}

func checkRowsAffected(taskResults sql.Result) (bool, error) {
	rows, err := taskResults.RowsAffected()
	if err != nil {
		return false, err
	}
	if rows == 0 {
		return false, nil
	}
	return true, nil
}

func GetTaskID(taskDB *sql.DB, taskTitle string) (int, error) {
	var taskID int
	err := taskDB.QueryRow("SELECT id FROM tasks WHERE task_title = (?)", taskTitle).Scan(&taskID)
	if errors.Is(err, sql.ErrNoRows) {
		return taskID, ErrTaskNotFound
	}
	if err != nil {
		return taskID, err
	}
	return taskID, nil
}
