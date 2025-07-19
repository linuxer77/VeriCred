package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/models"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUser function called.")

	var user models.Users
	
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	result := db.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	log.Println("Successfully created the user...")
	log.Println("result: ", result)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

