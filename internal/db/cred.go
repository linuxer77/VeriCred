package db
import (
	"errors"
	"gorm.io/gorm"
	"vericred/internal/models"
)

func StoreCredential(cred models.Credential) error {
    result := DB.Create(&cred)
    if result.Error != nil {
        return result.Error
    }
    return nil
}
	

// GetCredential retrieves a credential by ID
func GetCredential(credentialID string) (*models.Credential, error) {
    var credential models.Credential
    result := DB.Where("id = ?", credentialID).First(&credential)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("credential not found")
        }
        return nil, result.Error
    }
    return &credential, nil
}

// GetUniversity retrieves university information
func GetUniversity(universityID string) (*models.University, error) {
    var university models.University
    result := DB.Where("id = ?", universityID).First(&university)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("university not found")
        }
        return nil, result.Error
    }
    return &university, nil
}

// Additional CRUD functions using models.Credential instead of db.Credential
func GetCredentialsByStudent(studentID string) ([]models.Credential, error) {
    var credentials []models.Credential
    result := DB.Where("student_id = ?", studentID).Find(&credentials)
    if result.Error != nil {
        return nil, result.Error
    }
    return credentials, nil
}

// GetCredentialsByUniversity retrieves all credentials issued by a university
func GetCredentialsByUniversity(universityID string) ([]models.Credential, error) {
    var credentials []models.Credential
    result := DB.Where("university_id = ?", universityID).Find(&credentials)
    if result.Error != nil {
        return nil, result.Error
    }
    return credentials, nil
}

// UpdateCredentialStatus updates the status of a credential
func UpdateCredentialStatus(credentialID string, status string) error {
    result := DB.Model(&models.Credential{}).Where("id = ?", credentialID).Update("status", status)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("credential not found")
    }
    return nil
}

// StoreUniversity saves university information
func StoreUniversity(univ models.University) error {
    result := DB.Create(&univ)
    if result.Error != nil {
        return result.Error
    }
    return nil
}