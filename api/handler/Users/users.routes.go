package users

import (
	"github.com/gorilla/mux"

	user_controller "github.com/zeroprograming/go-postgresql-rest-api/api/handler/Users/controller"
)

// UsersRoutes configures the routes and assigns the corresponding handlers for user-related operations.
func UsersRoutes(router *mux.Router) {
	// Configure routes and assign corresponding handlers
	router.HandleFunc("/users", user_controller.GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", user_controller.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", user_controller.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user", user_controller.DeleteUserHandler).Methods("DELETE")
}
