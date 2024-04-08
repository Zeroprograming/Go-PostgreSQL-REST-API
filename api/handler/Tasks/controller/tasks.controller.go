package tasks_controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeroprograming/go-postgresql-rest-api/internal/database"
	"github.com/zeroprograming/go-postgresql-rest-api/internal/database/models"
)

// GetTasksHandler retrieves all tasks from the database and returns them as JSON.
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve all tasks from the database
	tasks := []models.Task{}
	database.DB.Find(&tasks)
	// Serialize the response
	response, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the tasks array"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetTaskHandler retrieves a specific task by its ID and returns it as JSON.
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Get the task ID from the URL parameters
	params := mux.Vars(r)

	// Find the task by ID
	task := models.Task{}
	result := database.DB.Unscoped().First(&task, "id = ?", params["id"])
	if result.Error != nil {
		// If there is an error while finding the task, return a 404 (Not Found) error
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Task not found"}`))
		return
	}
	// Serialize the response
	response, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the task"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateTaskHandler creates a new task and stores it in the database.
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	task := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error decoding the request body"}`))
		return
	}

	// Check if the user associated with the task exists
	user := database.DB.First(&models.User{}, task.UserId)
	if user.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	// Create the task in the database
	database.DB.Create(&task)

	// Serialize the response
	response, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the task"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// UpdateTaskHandler updates an existing task in the database.
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Get the task ID from the URL parameters
	params := mux.Vars(r)

	// Find the task by ID
	task := models.Task{}
	result := database.DB.Unscoped().First(&task, "id = ?", params["id"])
	if result.Error != nil {
		// If there is an error while finding the task, return a 404 (Not Found) error
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Task not found"}`))
		return
	}

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error decoding the request body"}`))
		return
	}

	// Update the task in the database
	database.DB.Save(&task)

	// Serialize the response
	response, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the task"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// DeleteTaskHandler deletes an existing task from the database.
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Get the task ID from the URL parameters
	params := mux.Vars(r)

	// Find the task by ID
	task := models.Task{}
	result := database.DB.Unscoped().First(&task, "id = ?", params["id"])
	if result.Error != nil {
		// If there is an error while finding the task, return a 404 (Not Found) error
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Task not found"}`))
		return
	}

	// Delete the task from the database
	database.DB.Unscoped().Delete(&task)

	w.WriteHeader(http.StatusNoContent)
}
