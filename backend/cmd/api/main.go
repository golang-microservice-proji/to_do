package main

import (
	"fmt"
	"net/http"
	"to-do-list/backend/api/handlers"
	"to-do-list/backend/config"
	"to-do-list/backend/pkg/task"
	"to-do-list/backend/pkg/timer"
	"to-do-list/backend/utils/alert"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	taskService := task.NewService()
	alertService := alert.NewAlertService()
	timerService := timer.NewService(alertService)

	taskHandler := handlers.NewTaskHandler(taskService, timerService)

	router := mux.NewRouter()

	router.HandleFunc("/tasks", taskHandler.HandleTasks).Methods("POST", "GET")
	router.HandleFunc("/tasks/{id}", taskHandler.HandleTask).Methods("GET", "PUT", "DELETE")

	fmt.Printf("Starting server on port %s...\n", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
