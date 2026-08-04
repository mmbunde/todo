package models

import (
	"database/sql"
	"strings"
	"testing"

	_ "modernc.org/sqlite"
)

func TestAddTask(t *testing.T) {
	testDB := InitTestDB(t)
	err := AddTask(testDB, "test task")
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}
}

func TestDupAddTask(t *testing.T) {
	testDB := InitTestDB(t)
	_ = AddTask(testDB, "test task")
	err := AddTask(testDB, "test task")
	if !strings.Contains(err.Error(), "Task already exists") {
		t.Fatalf("Failed duplicate task add: %v", err)
	}
}

// Need to isolate ListTasks in future. This is intergration testing right now not a unit test
func TestListTasks(t *testing.T) {
	testDB := InitTestDB(t)
	_ = AddTask(testDB, "test task")
	task, err := ListTask(testDB)
	if err != nil {
		t.Fatalf("Failed to list tasks: %v", err)
	}
	if len(task) != 1 {
		t.Fatalf("Tasks not the right number of tasks to list")
	}
	if task[0].ID != 1 {
		t.Fatalf("Incorrect task ID")
	}
	if task[0].TaskTitle != "test task" {
		t.Fatal("Incorrect task title")
	}
	if task[0].Complete != 0 {
		t.Fatalf("Incorrect default complete value")
	}
}

func TestEmptyList(t *testing.T) {
	testDB := InitTestDB(t)
	tasks, err := ListTask(testDB)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(tasks) != 0 {
		t.Fatalf("Expected an empty list but got a length of %v", len(tasks))
	}
}

// Integration test with AddTask and ListTask. Need to work on unit tests
func TestCompleteTaskSuccess(t *testing.T) {
	testDB := InitTestDB(t)
	_ = AddTask(testDB, "test task")
	err := CompleteTask(testDB, "test task")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	task, _ := ListTask(testDB)
	if task[0].Complete != 1 {
		t.Fatalf("Unable to complete task: %v", err)
	}
}

func TestCompleteTaskNotFound(t *testing.T) {
	testDB := InitTestDB(t)
	err := CompleteTask(testDB, "test task")
	if err == nil {
		t.Fatalf("Error equal to nil. Completed a task: %v", err)
	}
	if !strings.Contains(err.Error(), "Task doesn't exist") {
		t.Fatalf("Expected no task to be found but one was found: %v", err)
	}
}

func TestCompleteTaskAlreadyComplete(t *testing.T) {
	testDB := InitTestDB(t)
	_ = AddTask(testDB, "test task")
	_ = CompleteTask(testDB, "test task")
	err := CompleteTask(testDB, "test task")
	if !strings.Contains(err.Error(), "Task is already complete") {
		t.Fatalf("Completed task doesn't show as complete: %v", err)
	}
}

func TestDeleteTaskSuccess(t *testing.T) {
	testDB := InitTestDB(t)
	_ = AddTask(testDB, "test task")
	err := DeleteTask(testDB, "test task")
	if err != nil {
		t.Fatalf("Unable to delete task: %v", err)
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	testDB := InitTestDB(t)
	err := DeleteTask(testDB, "test task")
	if err == nil {
		t.Fatalf("Error was nil and was able to delete a task: %v", err)
	}
	if !strings.Contains(err.Error(), "No task found to delete") {
		t.Fatalf("Expected no task to be deleted but one wass: %v", err)
	}
}

func InitTestDB(t *testing.T) *sql.DB {
	testDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test db: %v", err)
	}
	_, err = testDB.Exec(`CREATE TABLE tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	task_title TEXT UNIQUE,
	complete INTEGER DEFAULT 0)`)
	if err != nil {
		t.Fatalf("Failed to open test db: %v", err)
	}
	return testDB
}
