package api

import (
	"net/http"
)

func SetupRoutes(handler APIHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", handler.ListTask)
	mux.HandleFunc("POST /tasks/{task_title}", handler.AddTask)
	return mux
}
