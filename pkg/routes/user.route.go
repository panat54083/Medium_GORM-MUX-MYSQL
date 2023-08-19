package routes

import (
	"example/backend/pkg/controllers"

	"github.com/gorilla/mux"
)

var InitUserRoute = func(router *mux.Router) {
	router.HandleFunc("/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/{id}", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteUserById).Methods("DELETE")

}
