package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/middleware"
	"vericred/internal/models"

	"gorm.io/gorm"
)


func CreateUniversity(w http.ResponseWriter, r *http.Request) {

	var body map[string]any
    err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(body)

	acad_email, ok := body["AcadEmail"].(string)
	if !ok {
		http.Error(w, "Invalid or missing AcadEmail", http.StatusBadRequest)
		return
	}
	org_name, ok := body["OrgName"].(string)
	if !ok {
		http.Error(w, "Invalid or missing OrgName", http.StatusBadRequest)
		return
	}
	org_type, ok := body["OrgType"].(string)
	if !ok {
		http.Error(w, "Invalid or missing OrgType", http.StatusBadRequest)
		return
	}
	org_url, ok := body["OrgUrl"].(string)
	if !ok {
		http.Error(w, "Invalid or missing OrgUrl", http.StatusBadRequest)
		return
	}
	org_desc, ok := body["OrgDesc"].(string)
	if !ok {
		http.Error(w, "Invalid or missing OrgDesc", http.StatusBadRequest)
		return
	}
	country, ok := body["Country"].(string)
	if !ok {
		http.Error(w, "Invalid or missing Country", http.StatusBadRequest)
		return
	}
	state, ok := body["State"].(string)
	if !ok {
		http.Error(w, "Invalid or missing State", http.StatusBadRequest)
		return
	}
	city, ok := body["City"].(string)
	if !ok {
		http.Error(w, "Invalid or missing City", http.StatusBadRequest)
		return
	}
	total_students, ok := body["TotalStudents"]
	if !ok {
		http.Error(w, "Invalid or missing TotalStudents", http.StatusBadRequest)
		return
	}
	students := total_students.(int)
	address, ok := body["Address"].(string)
	if !ok {
		http.Error(w, "Invalid or missing Address", http.StatusBadRequest)
		return
	}
	postal_code, ok := body["PostalCode"].(string)
	if !ok {
		http.Error(w, "Invalid or missing PostalCode", http.StatusBadRequest)
		return
	}


	metamaskAddress, ok := r.Context().Value(middleware.MetamaskAddressKey).(string)
	if !ok || metamaskAddress == "" {
		http.Error(w, "metamaskAddress is missing or invalid", http.StatusBadRequest)
		return
	}

	org := models.Organization{
		MetamaskAddress: metamaskAddress,
		AcadEmail: acad_email,
		OrgName: org_name,
		OrgType: org_type,
		OrgUrl: org_url,
		OrgDesc: org_desc,
		Country: country,
		State: state,
		City: city,
		TotalStudents: students,
		Address: address,
		PostalCode: postal_code,
	}

	
	res := db.DB.Where("metamask_address = ?", metamaskAddress).First(&org)

	if res.Error == gorm.ErrRecordNotFound {
		createResult := db.DB.Create(&org)
        if createResult.Error != nil {
            http.Error(w, "Failed to create account", http.StatusInternalServerError)
            return
        }
		
		var account models.Accounts

		if err := db.DB.Where("metamask_address = ?", metamaskAddress).First(&account).Error; err != nil {
			log.Fatalf("account not found: %v", err)
		}

		account.OwnerID = org.ID
		account.OwnerType = "university"

		if err := db.DB.Save(&account).Error; err != nil {
			log.Fatalf("Failed to update account: %v", err)
		}
	
	} else if res.Error != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org)
}

// func ShowUser(w http.ResponseWriter, r *http.Request) {
// 	metamaskAddress := r.Context().Value("metamaskAddress").(string)
// 	var user models.Users
// 	db.DB.First(&user, metamaskAddress)
// 	json.NewEncoder(w).Encode(user)
// }