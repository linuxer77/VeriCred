package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/models"
	"vericred/pkg"
	"vericred/redisdb"

	"gorm.io/gorm"
)


func LoginInMetamask(w http.ResponseWriter, r *http.Request) {
    var body struct {
        MetamaskAddress string `json:"metamask_addresse"`
        Signature       string `json:"signature"`
    }
    
    json.NewDecoder(r.Body).Decode(&body)

    var n redisdb.Redis
    nonce := n.RedisGetNonce(body.MetamaskAddress)

    log.Println("Nonce: ", nonce)
    
    if nonce == "" {
        http.Error(w, "nonce missing", http.StatusUnauthorized)
        return
    }
    
    log.Println("Nonce: ", nonce)

    verified, err := pkg.VerifySignature(body.MetamaskAddress, nonce, body.Signature)
    
    if err != nil || !verified {
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }
        
    var acc models.Accounts
    err = json.NewDecoder(r.Body).Decode(&acc)
    if err != nil {
        http.Error(w, "Invalid JSON request", http.StatusBadRequest)
        return
    }
    
    var existingAcc models.Accounts
    result := db.DB.First(&existingAcc, "metamask_address = ?", acc.MetamaskAddress)
      
    if result.Error == gorm.ErrRecordNotFound {
        createResult := db.DB.Create(&acc)
        if createResult.Error != nil {
            http.Error(w, "Failed to create account", http.StatusInternalServerError)
            return
        }
    } else if result.Error != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    tokenStr, err := pkg.CreateToken(acc.MetamaskAddress)
    if err != nil {
        http.Error(w, "Failed to create token", http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, tokenStr)
}