package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/models"
	"vericred/pkg"

	"gorm.io/gorm"
)


func LoginInMetamask(w http.ResponseWriter, r *http.Request) {
    var acc models.Accounts
    err := json.NewDecoder(r.Body).Decode(&acc)
    if err != nil {
        http.Error(w, "Invalid JSON request", http.StatusBadRequest)
        return
    }
    
    // Check if account exists
    var existingAcc models.Accounts
    result := db.DB.First(&existingAcc, "metamask_address = ?", acc.MetamaskAddress)
    
    if result.Error == gorm.ErrRecordNotFound {
        // New user - create account
        createResult := db.DB.Create(&acc)
        if createResult.Error != nil {
            http.Error(w, "Failed to create account", http.StatusInternalServerError)
            return
        }
    } else if result.Error != nil {
        // Database error
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    // If no error, user exists - proceed to create token
    
    // Create JWT token
    tokenStr, err := pkg.CreateToken(acc.MetamaskAddress)
    if err != nil {
        http.Error(w, "Failed to create token", http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, tokenStr)
}