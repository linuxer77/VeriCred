package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(db *gorm.DB) {
	DB = db
}

type Accounts struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	MetamaskAddress string     `gorm:"unique;not null;size:42" json:"metamask_address"`
	AccountType     string     `gorm:"not null" json:"account_type"`
	Verified        bool       `gorm:"default:false" json:"verified"`
	Credentials     int        `gorm:"default:0" json:"credentials"`
	LastLoginAt     *time.Time `json:"last_login_at"`
	IsActive        bool       `gorm:"default:true" json:"is_active"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type Users struct {
	MetamaskAddress     string `gorm:"primaryKey;size:42" json:"metamask_address"`
	Email               string `gorm:"not null;unique" json:"email"`
	FirstName           string `gorm:"not null" json:"first_name"`
	LastName            string `gorm:"not null" json:"last_name"`
	Password            string `gorm:"not null" json:"-"`
	StudentID           string `gorm:"size:100" json:"student_id"`
	CurrentUniversityID string `gorm:"size:100" json:"current_university_id"`
	ProfilePicture      string `gorm:"size:255" json:"profile_picture"`

	Account Accounts `gorm:"foreignKey:MetamaskAddress;references:MetamaskAddress"`
}

type Organization struct {
	MetamaskAddress string `gorm:"primaryKey;size:42" json:"metamask_address"`
	AcadEmail       string `gorm:"not null;unique" json:"acad_email"`
	OrgName         string `gorm:"not null" json:"org_name"`
	OrgType         string `gorm:"size:100" json:"org_type"`
	OrgUrl          string `gorm:"size:255" json:"org_url"`
	LogoIPFSHash    string `gorm:"size:100" json:"logo_ipfs_hash"`
	OrgDesc         string `gorm:"type:text" json:"org_desc"`
	Country         string `gorm:"size:100" json:"country"`
	State           string `gorm:"size:100" json:"state"`
	City            string `gorm:"size:100" json:"city"`
	Address         string `gorm:"type:text" json:"address"`
	PostalCode      string `gorm:"size:20" json:"postal_code"`
	PhoneNumber     string `gorm:"size:20" json:"phone_number"`
	EstablishedYear int    `json:"established_year"`
	Accreditation   string `gorm:"size:255" json:"accreditation"`
	IsVerified      bool   `gorm:"default:false" json:"is_verified"`

	TotalCredentials  int `gorm:"default:0" json:"total_credentials"`
	ActiveCredentials int `gorm:"default:0" json:"active_credentials"`
	TotalStudents     int `gorm:"default:0" json:"total_students"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Account Accounts `gorm:"foreignKey:MetamaskAddress;references:MetamaskAddress"`
}

type Degree struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    DegreeName  string `gorm:"not null;size:255" json:"degree_name"`
    DegreeType  string `gorm:"not null;size:100" json:"degree_type"` // Bachelor's, Master's, PhD, etc.
    Description string `gorm:"type:text" json:"description"`
    
    // University relationship
    UniversityWallet string       `gorm:"not null;size:42" json:"university_wallet"`
    University       Organization `gorm:"foreignKey:UniversityWallet;references:MetamaskAddress"`
    
}

type Credential struct {
	ID              string    `gorm:"primaryKey;size:100" json:"id"`
	DegreeID         uint      `gorm:"not null" json:"degree_id"`
    StudentWallet    string    `gorm:"not null;size:42" json:"student_wallet"`
    UniversityWallet string    `gorm:"not null;size:42" json:"university_wallet"`
    DegreeName      string    `gorm:"not null;size:255" json:"degree_name"`
	Description     string    `gorm:"type:text" json:"description"`
	Type            string    `gorm:"not null;size:100" json:"type"`
	Major           string    `gorm:"size:255" json:"major"`
	IssuedDate      time.Time `gorm:"not null" json:"issued_date"`
	GraduationDate  string    `gorm:"size:50" json:"graduation_date"`
	Status          string    `gorm:"default:Active;size:50;index" json:"status"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	
	Student    Users              `gorm:"foreignKey:StudentWallet;references:MetamaskAddress"`
	University Organization       `gorm:"foreignKey:UniversityWallet;references:MetamaskAddress"`
	Degree     Degree       `gorm:"foreignKey:DegreeID;references:ID"`
}

type StudentAcademicInfo struct {
	MetamaskAddress string `json:"metamask_address"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	StudentID       string `json:"student_id"`

	Major          string `json:"major"`
	GPA            string `json:"gpa"`
	GraduationDate string `json:"graduation_date"`
	Status         string `json:"status"`

	DegreeType     string `json:"degree_type"`
	DegreeTitle    string `json:"degree_title"`
	CourseDuration string `json:"course_duration"`

	TotalCredentials int    `json:"total_credentials"`
	EnrollmentStatus string `json:"enrollment_status"`
}

func GetUniversityStudents(universityWallet string) ([]Users, error) {
	var students []Users

	result := DB.Table("users").
		Joins("JOIN credentials ON users.metamask_address = credentials.student_wallet").
		Where("credentials.university_wallet = ?", universityWallet).
		Group("users.metamask_address").
		Find(&students)

	return students, result.Error
}

func GetStudentAcademicInfo(universityWallet string) ([]StudentAcademicInfo, error) {
	var students []StudentAcademicInfo

	result := DB.Table("users u").
		Select(`
            u.metamask_address,
            u.first_name,
            u.last_name,
            u.email,
            u.student_id,
            c.major,
            c.gpa,
            c.graduation_date,
            c.status,
            ct.degree_type,
            ct.department,
            ct.course_duration,
            COUNT(c.id) as total_credentials,
            CASE 
                WHEN c.status = 'Active' AND c.graduation_date <= NOW() THEN 'graduated'
                WHEN c.status = 'Active' THEN 'enrolled'
                ELSE 'inactive'
            END as enrollment_status
        `).
		Joins("JOIN credentials c ON u.metamask_address = c.student_wallet").
		Joins("LEFT JOIN credential_templates ct ON c.template_id = ct.id").
		Where("c.university_wallet = ?", universityWallet).
		Group("u.metamask_address, c.major, c.gpa, c.graduation_date, c.status, ct.degree_type, ct.department, ct.course_duration").
		Find(&students)

	return students, result.Error
}

type VerificationRequest struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	EmployerWallet string     `gorm:"not null;size:42" json:"employer_wallet"`
	StudentWallet  string     `gorm:"not null;size:42" json:"student_wallet"`
	CredentialID   string     `gorm:"not null" json:"credential_id"`
	RequestMessage string     `gorm:"type:text" json:"request_message"`
	Status         string     `gorm:"default:pending" json:"status"`
	RequestedAt    time.Time  `gorm:"autoCreateTime" json:"requested_at"`
	RespondedAt    *time.Time `json:"responded_at"`

	Student    Users      `gorm:"foreignKey:StudentWallet;references:MetamaskAddress"`
	Credential Credential `gorm:"foreignKey:CredentialID;references:ID"`
}

type UniversityDashboardStats struct {
	TotalStudents      int64 `json:"total_students"`
	ActiveCredentials  int64 `json:"active_credentials"`
	TotalCredentials   int64 `json:"total_credentials"`
	GraduatedStudents  int64 `json:"graduated_students"`
	PendingCredentials int64 `json:"pending_credentials"`

	DepartmentStats []DepartmentStat `json:"department_stats"`

	DegreeTypeStats []DegreeTypeStat `json:"degree_type_stats"`
}

type DepartmentStat struct {
	Department      string `json:"department"`
	StudentCount    int    `json:"student_count"`
	CredentialCount int    `json:"credential_count"`
}

type DegreeTypeStat struct {
	DegreeType   string `json:"degree_type"`
	StudentCount int    `json:"student_count"`
}

func GetUniversityDashboardStats(universityWallet string) (*UniversityDashboardStats, error) {
	stats := &UniversityDashboardStats{}

	DB.Table("credentials").
		Where("university_wallet = ?", universityWallet).
		Distinct("student_wallet").
		Count(&stats.TotalStudents)

	DB.Table("credentials").
		Where("university_wallet = ? AND status = ?", universityWallet, "Active").
		Count(&stats.ActiveCredentials)

	DB.Table("credentials").
		Where("university_wallet = ?", universityWallet).
		Count(&stats.TotalCredentials)

	DB.Table("credentials").
		Where("university_wallet = ? AND graduation_date <= NOW() AND status = ?", universityWallet, "Active").
		Distinct("student_wallet").
		Count(&stats.GraduatedStudents)

	DB.Table("credentials c").
		Select("ct.department, COUNT(DISTINCT c.student_wallet) as student_count, COUNT(c.id) as credential_count").
		Joins("JOIN credential_templates ct ON c.template_id = ct.id").
		Where("c.university_wallet = ?", universityWallet).
		Group("ct.department").
		Find(&stats.DepartmentStats)

	DB.Table("credentials c").
		Select("ct.degree_type, COUNT(DISTINCT c.student_wallet) as student_count").
		Joins("JOIN credential_templates ct ON c.template_id = ct.id").
		Where("c.university_wallet = ?", universityWallet).
		Group("ct.degree_type").
		Find(&stats.DegreeTypeStats)

	return stats, nil
}

type UserDashboardProfile struct {
	MetamaskAddress string `json:"metamask_address"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	StudentID       string `json:"student_id"`
	ProfilePicture  string `json:"profile_picture"`
	Bio             string `json:"bio"`
	PhoneNumber     string `json:"phone_number"`

	AccountType      string     `json:"account_type"`
	Verified         bool       `json:"verified"`
	TotalCredentials int        `json:"total_credentials"`
	LastLoginAt      *time.Time `json:"last_login_at"`

	UniversitiesAttended []string `json:"universities_attended"`
	LatestDegree         string   `json:"latest_degree"`
	LatestGPA            string   `json:"latest_gpa"`
	CredentialCount      int64    `json:"credential_count"`
}

/*func GetUserDashboardProfile(userWallet string) (*UserDashboardProfile, error) {
	profile := &UserDashboardProfile{}

	var user Users
	if err := DB.Where("metamask_address = ?", userWallet).First(&user).Error; err != nil {
		return nil, err
	}

	var account Accounts
	DB.Where("metamask_address = ?", userWallet).First(&account)

	profile.MetamaskAddress = user.MetamaskAddress
	profile.Email = user.Email
	profile.FirstName = user.FirstName
	profile.LastName = user.LastName
	profile.StudentID = user.StudentID
	profile.ProfilePicture = user.ProfilePicture

	profile.AccountType = account.AccountType
	profile.Verified = account.Verified
	profile.TotalCredentials = account.Credentials
	profile.LastLoginAt = account.LastLoginAt

	var credentials []Credential
	DB.Where("student_wallet = ?", userWallet).Find(&credentials)

	profile.CredentialCount = int64(len(credentials))

	universities := make(map[string]bool)
	for _, cred := range credentials {
		var org Organization
		if DB.Where("metamask_address = ?", cred.UniversityID).First(&org).Error == nil {
			universities[org.OrgName] = true
		}
	}

	for uni := range universities {
		profile.UniversitiesAttended = append(profile.UniversitiesAttended, uni)
	}

	if len(credentials) > 0 {
		latest := credentials[len(credentials)-1]
		profile.LatestDegree = latest.DegreeName
		profile.LatestGPA = latest.GPA
	}

	return profile, nil
}
*/