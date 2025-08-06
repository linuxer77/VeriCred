package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func GetRecentMintedCredentials(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    universityWallet := r.URL.Query().Get("university_wallet")
    if universityWallet == "" {
        http.Error(w, "University wallet required", http.StatusBadRequest)
        return
    }

    limitStr := r.URL.Query().Get("limit")
    limit := 10
    if limitStr != "" {
        if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
            limit = parsedLimit
        }
    }

    var credentials []models.Credential
    if err := models.DB.Where("university_wallet = ?", universityWallet).
        Order("created_at DESC").
        Limit(limit).
        Preload("Student").
        Find(&credentials).Error; err != nil {
        http.Error(w, "Failed to fetch recent credentials", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(credentials)
}

func GetMintRequests(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var requests []models.Credential
    if err := models.DB.Where("status = ?", "pending").Find(&requests).Error; err != nil {
        http.Error(w, "Failed to fetch mint requests", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(requests)
}


func SearchStudents(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    universityWallet := r.URL.Query().Get("university_wallet")
    if universityWallet == "" {
        http.Error(w, "University wallet required", http.StatusBadRequest)
        return
    }

    searchQuery := r.URL.Query().Get("search")
    if searchQuery == "" {
        http.Error(w, "Search query required", http.StatusBadRequest)
        return
    }

    limitStr := r.URL.Query().Get("limit")
    limit := 20
    if limitStr != "" {
        if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
            limit = parsedLimit
        }
    }

    var students []models.Users
    query := models.DB.Table("users").
        Select("users.*, credentials.university_wallet").
        Joins("JOIN credentials ON users.metamask_address = credentials.student_wallet").
        Where("credentials.university_wallet = ?", universityWallet).
        Where("users.first_name ILIKE ? OR users.last_name ILIKE ? OR users.student_id ILIKE ? OR users.metamask_address ILIKE ?", 
            "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").
        Group("users.metamask_address, credentials.university_wallet").
        Limit(limit)

    if err := query.Find(&students).Error; err != nil {
        http.Error(w, "Failed to search students", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(students)
}

func GetStudentsByFilter(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    universityWallet := r.URL.Query().Get("university_wallet")
    if universityWallet == "" {
        http.Error(w, "University wallet required", http.StatusBadRequest)
        return
    }

    mintingStatus := r.URL.Query().Get("minting") 
    
    var students []struct {
        models.Users
        CredentialCount   int    `json:"credential_count"`
        LatestStatus      string `json:"latest_status"`
        MintingStatus     string `json:"minting_status"`
        LatestDegree      string `json:"latest_degree"`
        GraduationDate    string `json:"graduation_date"`
    }

    query := models.DB.Table("users").
        Select(`users.*, 
                COUNT(credentials.id) as credential_count,
                MAX(credentials.status) as latest_status,
                CASE 
                    WHEN COUNT(CASE WHEN credentials.status = 'Active' THEN 1 END) > 0 THEN 'Minted'
                    WHEN COUNT(CASE WHEN credentials.status = 'Pending' THEN 1 END) > 0 THEN 'Pending'
                    ELSE 'No Requests'
                END as minting_status,
                MAX(credentials.degree_name) as latest_degree,
                MAX(credentials.graduation_date) as graduation_date`).
        Joins("LEFT JOIN credentials ON users.metamask_address = credentials.student_wallet AND credentials.university_wallet = ?", universityWallet).
        Group("users.metamask_address")

    if mintingStatus != "" {
    switch mintingStatus {
    case "minted":
        query = query.Having("COUNT(CASE WHEN credentials.status = 'Active' THEN 1 END) > 0")
    case "pending":
        query = query.Having("COUNT(CASE WHEN credentials.status = 'Pending' THEN 1 END) > 0")
    case "no_requests":
        query = query.Having("COUNT(credentials.id) = 0")
    }
}

    if err := query.Find(&students).Error; err != nil {
        http.Error(w, "Failed to fetch filtered students", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(students)
}

func GetUniversityDashboardStats(universityWallet string) (*models.UniversityDashboardStats, error) {
    stats := &models.UniversityDashboardStats{}

    models.DB.Table("credentials").
        Where("university_wallet = ?", universityWallet).
        Distinct("student_wallet").
        Count(&stats.TotalStudents)

    models.DB.Table("credentials").
        Where("university_wallet = ? AND status = ?", universityWallet, "Active").
        Count(&stats.ActiveCredentials)

    models.DB.Table("credentials").
        Where("university_wallet = ?", universityWallet).
        Count(&stats.TotalCredentials)

    models.DB.Table("credentials").
        Where("university_wallet = ? AND status = ?", universityWallet, "Active").
        Distinct("student_wallet").
        Count(&stats.GraduatedStudents)

    models.DB.Table("credentials").
        Where("university_wallet = ? AND status = ?", universityWallet, "Pending").
        Count(&stats.PendingCredentials)
    models.DB.Table("credentials").
        Select("major as department, COUNT(DISTINCT student_wallet) as student_count, COUNT(id) as credential_count").
        Where("university_wallet = ? AND major != ''", universityWallet).
        Group("major").
        Find(&stats.DepartmentStats)
    models.DB.Table("credentials c").
        Select("d.degree_type, COUNT(DISTINCT c.student_wallet) as student_count").
        Joins("JOIN degrees d ON c.degree_id = d.id").
        Where("c.university_wallet = ?", universityWallet).
        Group("d.degree_type").
        Find(&stats.DegreeTypeStats)

    return stats, nil
}

func GetUniversityDashboard(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    universityWallet := r.URL.Query().Get("university_wallet")
    if universityWallet == "" {
        http.Error(w, "University wallet required", http.StatusBadRequest)
        return
    }

    stats, err := GetUniversityDashboardStats(universityWallet)
    if err != nil {
        http.Error(w, "Failed to fetch dashboard stats", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(stats)
}

func UniversitySettings(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    universityWallet := r.URL.Query().Get("university_wallet")
    if universityWallet == "" {
        http.Error(w, "University wallet required", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case "GET":
        var university models.Organization
        if err := models.DB.Where("metamask_address = ?", universityWallet).First(&university).Error; err != nil {
            http.Error(w, "University not found", http.StatusNotFound)
            return
        }

        response := map[string]interface{}{
            "org_name":    university.OrgName,
            "org_desc":    university.OrgDesc,
            "org_url":     university.OrgUrl,
            "country":     university.Country,
            "state":       university.State,
            "city":        university.City,
            "phone":       university.PhoneNumber,
            "Academic_Email": university.AcadEmail,
        }

        json.NewEncoder(w).Encode(response)

    case "PUT":
        var updateData struct {
            OrgName string `json:"org_name"`
            OrgDesc string `json:"org_desc"`
            OrgUrl  string `json:"org_url"`
            Country string `json:"country"`
            State   string `json:"state"`
            City    string `json:"city"`
            Phone   string `json:"phone"`
            AcadEmail string `json:"academic_email"`
        }

        if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
            http.Error(w, "Invalid JSON format", http.StatusBadRequest)
            return
        }

        updates := map[string]interface{}{}
        if updateData.OrgName != "" {
            updates["org_name"] = updateData.OrgName
        }
        if updateData.OrgDesc != "" {
            updates["org_desc"] = updateData.OrgDesc
        }
        if updateData.OrgUrl != "" {
            updates["org_url"] = updateData.OrgUrl
        }
        if updateData.Country != "" {
            updates["country"] = updateData.Country
        }
        if updateData.State != "" {
            updates["state"] = updateData.State
        }
        if updateData.City != "" {
            updates["city"] = updateData.City
        }
        if updateData.Phone != "" {
            updates["phone_number"] = updateData.Phone
        }
        if updateData.AcadEmail != "" {
            updates["academic_email"] = updateData.AcadEmail
        }

        if err := models.DB.Model(&models.Organization{}).
            Where("metamask_address = ?", universityWallet).
            Updates(updates).Error; err != nil {
            http.Error(w, "Failed to update university settings", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{
            "message": "University settings updated successfully",
        })

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}