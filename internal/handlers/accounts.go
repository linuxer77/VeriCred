package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/models"
)


func CreateAccounts(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateAccounts funciton called.")

	var accounts models.Accounts

	err := json.NewDecoder(r.Body).Decode(&accounts)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&accounts)
	if result.Error != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	log.Println("Successfully created the account...")
	log.Println("result: ", result)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(accounts)


}