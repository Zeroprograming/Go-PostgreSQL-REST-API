package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	routes "github.com/zeroprograming/go-postgresql-rest-api/api/handler"
	"github.com/zeroprograming/go-postgresql-rest-api/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.DBConnection()

	router := mux.NewRouter()
	log.Println("Server running on port", "http://localhost"+os.Getenv("PORT"))

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(os.Getenv("PORT"), router)

}
