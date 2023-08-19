package controllers

import (
	"encoding/json"
	"example/backend/pkg/config"
	"example/backend/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Get the request body and decode it into a new user object
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the new user in the database
	config.GetDB().Create(&newUser)

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the newly created user as JSON and send it as the response
	err := json.NewEncoder(w).Encode(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Define a slice to store all user records
	var allUsers []*models.User

	// Retrieve all user records from the database
	config.GetDB().Find(&allUsers)

	// Set the response headers to specify JSON format
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)

	// Encode the list of users as JSON and send it as the response
	err := json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Initialize a user variable to store the user record
	var user models.User

	// Retrieve the user ID from the URL parameters
	id := mux.Vars(r)["id"]

	// Query the database to find the user by their ID
	config.GetDB().First(&user, id)

	// Check if the user was not found (user.ID == 0)
	if user.ID == 0 {
		// Encode and send a response indicating that the user was not found
		json.NewEncoder(w).Encode("User not found!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the user record as JSON and send it as the response
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := mux.Vars(r)["id"]
	db := config.GetDB()

	db.First(&user, id)

	//Check if null
	if user.ID == 0 {
		json.NewEncoder(w).Encode("user not found!")
		return
	}

	// Get new User info from request
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)

	// Send new User info as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := mux.Vars(r)["id"]
	db := config.GetDB()
	db.First(&user, id)
	if user.ID == 0 {
		json.NewEncoder(w).Encode("user not found!")
		return
	}
	db.Delete(&user, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("user deleted successfully")
}
