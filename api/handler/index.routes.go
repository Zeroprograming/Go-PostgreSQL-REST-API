package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	tasks "github.com/zeroprograming/go-postgresql-rest-api/api/handler/Tasks"
	users "github.com/zeroprograming/go-postgresql-rest-api/api/handler/Users"
)

// ApiRoutes configures the API routes for the application.
func ApiRoutes(router *mux.Router) {
	// Define the root route for the ping handler
	router.HandleFunc("/", PingHandler).Methods("GET")

	// Configure routes for users and tasks
	users.UsersRoutes(router)
	tasks.TasksRoutes(router)
}

// PingHandler handles the ping request to check if the server is up.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map for the response JSON
	response := map[string]interface{}{
		"status": http.StatusOK,
		"data":   "pong",
	}

	// Convert the map to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the content type and write the JSON response to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
