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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Invalid JSON request")
		return
	}
	
	tempAcc := db.DB.First(&acc, "MetamaskAddress = ?", acc.MetamaskAddress)
	
	if (tempAcc.Error == gorm.ErrRecordNotFound) && (tempAcc.RowsAffected == 0) {
		result := db.DB.Create(&acc)
		if result.Error != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}
		tokenStr, err := pkg.CreateToken(acc.MetamaskAddress)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error when creating token: ", err)
		}
		w.WriteHeader(http.StatusOK)
	    fmt.Fprint(w, tokenStr)
		return
	} else if (tempAcc.Error != nil) && (tempAcc.RowsAffected == 1) {
		tokenStr, err := pkg.CreateToken(acc.MetamaskAddress)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error when creating token: ", err)
		}
		w.WriteHeader(http.StatusOK)
	    fmt.Fprint(w, tokenStr)
	}
	
}

