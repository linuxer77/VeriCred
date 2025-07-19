package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vericred/internal/models"
	"vericred/pkg"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u models.Users
	json.NewDecoder(r.Body).Decode(&u)

	// TODO: implement the real way of verifying im hard coding this shit for now
	if (u.MetamaskAddress == "xyz") {
		tokenString, err := pkg.CreateToken(u.MetamaskAddress)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error when creating token: ", err)
		}
		w.WriteHeader(http.StatusOK)
	    fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	    fmt.Fprint(w, "Invalid credentials")
	}
}