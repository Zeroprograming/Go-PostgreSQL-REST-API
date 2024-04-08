package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	routes "github.com/zeroprograming/go-postgresql-rest-api/api/handler"
	"github.com/zeroprograming/go-postgresql-rest-api/internal/database"
)

// main function is the entry point of the application
func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	database.DBConnection()

	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Use logging middleware for debugging
	router.Use(LoggingMiddleware)

	// Get the port from environment variables
	port := os.Getenv("PORT")
	log.Println("Server running on port", "http://localhost"+port)

	// Configure API routes
	routes.ApiRoutes(router)

	// Start the HTTP server
	http.ListenAndServe(port, router)

}

// LoggingMiddleware is a middleware function for logging HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
