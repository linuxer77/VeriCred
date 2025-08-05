package handlers
import (
	"encoding/json"
	"net/http"
	"vericred/internal/models"
)

func GetStudentCredentials(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    studentWallet := r.URL.Query().Get("wallet")
    if studentWallet == "" {
        http.Error(w, "Student wallet required", http.StatusBadRequest)
        return
    }

    var credentials []models.Credential
    if err := models.DB.Where("student_wallet = ?", studentWallet).Find(&credentials).Error; err != nil {
        http.Error(w, "Failed to fetch credentials", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(credentials)
}