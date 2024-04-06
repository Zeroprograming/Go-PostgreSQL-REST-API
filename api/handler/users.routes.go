package routes

import (
	"encoding/json"
	"net/http"

	"github.com/zeroprograming/go-postgresql-rest-api/internal/database"
	"github.com/zeroprograming/go-postgresql-rest-api/internal/database/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener todos los usuarios
	users := []models.User{}
	database.DB.Find(&users)
	// Serializar la respuesta
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

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de la URL
	id := r.URL.Query().Get("id")
	// Buscar el usuario por ID
	user := models.User{}
	database.DB.First(&user, id)
	// Serializar la respuesta
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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error decoding the request body"}`))
		return
	}
	// Crear el usuario en la base de datos
	user.SetPassword(user.Password)
	database.DB.Create(&user)
	// Serializar la respuesta
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
