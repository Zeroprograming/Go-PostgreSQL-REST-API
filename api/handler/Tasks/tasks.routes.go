package tasks

import (
	"github.com/gorilla/mux"

	task_controller "github.com/zeroprograming/go-postgresql-rest-api/api/handler/Tasks/controller"
)

// TasksRoutes configures routes for task-related HTTP requests and assigns corresponding handlers.
func TasksRoutes(router *mux.Router) {
	// Configure routes and assign corresponding handlers
	router.HandleFunc("/tasks", task_controller.GetTasksHandler).Methods("GET")
	router.HandleFunc("/task/{id}", task_controller.GetTaskHandler).Methods("GET")
	router.HandleFunc("/task", task_controller.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/task/{id}", task_controller.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/task", task_controller.DeleteTaskHandler).Methods("DELETE")
}
