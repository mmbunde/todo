package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/mmbunde/todo/api"
	"github.com/mmbunde/todo/storage"
)

func main() {
	var taskDB string
	flag.StringVar(&taskDB, "f", "tasks", "The DB for your tasks")
	flag.Parse()

	db, err := storage.InitDB(taskDB)
	if err != nil {
		fmt.Println(err)
	}

	handler := api.NewAPIHandler(db)
	router := api.SetupRoutes(handler)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}

}
