package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/models"
)


func CreateOrg(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateOrg funciton called.")

	var org models.Organization

	err := json.NewDecoder(r.Body).Decode(&org)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&org)
	if result.Error != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	log.Println("Successfully created the org...")
	log.Println("result: ", result)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org)


}