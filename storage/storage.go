package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func SetFilePath(fileName string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filePath := filepath.Join(homeDir, ".config", "todo", fileName)
	return filePath
}

func InitDB(fileName string) (*sql.DB, error) {
	filePath := SetFilePath(fileName)
	taskDB, err := sql.Open("sqlite", filePath)
	if err != nil {
		return taskDB, err
	}
	_, err = taskDB.Exec(`CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	task_title TEXT UNIQUE,
	complete INTEGER DEFAULT 0)`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return taskDB, nil
}
