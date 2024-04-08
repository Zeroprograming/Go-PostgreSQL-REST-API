// Package user_controller provides HTTP request handlers for user-related operations.
package user_controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeroprograming/go-postgresql-rest-api/internal/database"
	"github.com/zeroprograming/go-postgresql-rest-api/internal/database/models"
)

// GetUsersHandler handles the HTTP GET request to retrieve all users.
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve all users from the database
	users := []models.User{}
	database.DB.Find(&users)

	// Serialize the response
	response, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the users array"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetUserHandler handles the HTTP GET request to retrieve a specific user by ID.
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameters
	params := mux.Vars(r)

	// Find the user by ID
	user := models.User{}
	result := database.DB.Unscoped().First(&user, "id = ?", params["id"])
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	// Serialize the response
	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the user"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateUserHandler handles the HTTP POST request to create a new user.
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body to get the user details
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error decoding the request body"}`))
		return
	}

	// Create the user in the database
	database.DB.Create(&user)

	// Serialize the response
	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the user"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// DeleteUserHandler handles the HTTP DELETE request to delete a user by ID.
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameters
	id := r.URL.Query().Get("id")

	// Find the user by ID
	user := models.User{}
	database.DB.First(&user, "id = ?", id)

	// Delete the user from the database
	database.DB.Delete(&user)

	// Serialize the response
	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the response"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
