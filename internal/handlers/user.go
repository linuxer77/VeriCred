package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"

	"gorm.io/gorm"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {

	var body map[string]any
    json.NewDecoder(r.Body).Decode(&body)

    email := body["email"].(string)
    first_name := body["firstName"].(string)
    last_name := body["lastName"].(string)
    student_id := body["studentEmail"].(string)
	
	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}

	user := models.Users{
		MetamaskAddress: metamaskAddress,
		Email: email,
		FirstName: first_name,
		LastName: last_name,
		StudentEmail: student_id,
	}

	
	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&user)

	if res.Error == gorm.ErrRecordNotFound {
		user.IsVerified = true
		createResult := db.DB.Create(&user)
        if createResult.Error != nil {
            http.Error(w, "Failed to create account", http.StatusInternalServerError)
            return
        }
		
		var account models.Accounts

		if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&account).Error; err != nil {
			log.Fatalf("account not found: %v", err)
		}

		account.OwnerID = user.ID
		account.OwnerType = "user"

		if err := db.DB.Save(&account).Error; err != nil {
			log.Fatalf("Failed to update account: %v", err)
		}
	
	} else if res.Error != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	metamaskAddress, ok := r.Context().Value("metamaskAddress").(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
	}
	log.Println("User logged in...")
	var user models.Users
	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}  else if res.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}