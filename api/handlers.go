package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mmbunde/todo/models"
)

type APIHandler struct {
	db *sql.DB
}

func NewAPIHandler(db *sql.DB) APIHandler {
	return APIHandler{db}
}

func (handler APIHandler) ListTask(writer http.ResponseWriter, request *http.Request) {
	tasks, err := models.ListTask(handler.db)
	if err != nil {
		log.Printf("error: list tasks: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = headerHelper(tasks, writer)
	if err != nil {
		log.Printf("error: encode task list, %v", err)
		return
	}
}

func (handler APIHandler) AddTask(writer http.ResponseWriter, request *http.Request) {
	task := request.PathValue("task_title")

	err := models.AddTask(handler.db, task)
	if err != nil {
		if errors.Is(err, models.ErrTaskExist) {
			log.Printf("error: add task, %v", err)
			http.Error(writer, err.Error(), http.StatusConflict)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = headerHelper(task, writer)
	if err != nil {
		log.Printf("error: encode add task, %v", err)
		return
	}
}

func (handler APIHandler) CompleteTask(writer http.ResponseWriter, request *http.Request) {
	task := request.PathValue("task_title")

	err := models.CompleteTask(handler.db, task)
	if err != nil {
		if errors.Is(err, models.ErrTaskNotFound) {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		} else if errors.Is(err, models.ErrTaskComplete) {
			http.Error(writer, err.Error(), http.StatusConflict)
			return
		} else if errors.Is(err, models.ErrTaskUpdateFailed) {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	err = headerHelper(task, writer)
	if err != nil {
		log.Printf("error: encode complete task, %v", err)
		return
	}
}

func (handler APIHandler) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	task := request.URL.Query().Get("task_title")

	err := models.DeleteTask(handler.db, task)
	if err != nil {
		if errors.Is(err, models.ErrTaskNotFound) {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = headerHelper(task, writer)
	if err != nil {
		log.Printf("error: encode delete task, %v", err)
	}

}

func headerHelper[T encodeable](task T, writer http.ResponseWriter) error {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(task)
	return err

}

type encodeable interface {
	string | []models.Task
}
