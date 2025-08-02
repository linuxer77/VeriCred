package models

import (
    "time"
    "gorm.io/datatypes"
    "gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(db *gorm.DB) {
    DB = db
}

type Accounts struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    MetamaskAddress string    `gorm:"unique;not null;size:42" json:"metamask_address"`
    AccountType     string    `gorm:"not null" json:"account_type"` // "user", "organization", "employer", "admin"
    Verified        bool      `gorm:"default:false" json:"verified"`
    Credentials     int       `gorm:"default:0" json:"credentials"` // Fixed typo
    LastLoginAt     *time.Time `json:"last_login_at"` // Track login activity
    IsActive        bool      `gorm:"default:true" json:"is_active"`
    CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Users struct {
    MetamaskAddress string    `gorm:"primaryKey;size:42" json:"metamask_address"` // Same as Accounts
    Email           string    `gorm:"not null;unique" json:"email"` // Fixed GORM tag
    FirstName       string    `gorm:"not null" json:"first_name"` // Fixed GORM tag
    LastName        string    `gorm:"not null" json:"last_name"` // Fixed GORM tag
    Password        string    `gorm:"not null" json:"-"` // Hidden from JSON
    StudentID       string    `gorm:"size:100" json:"student_id"` // University student ID
    CurrentUniversityID string `gorm:"size:100" json:"current_university_id"` // If currently enrolled
    ProfilePicture  string    `gorm:"size:255" json:"profile_picture"` // Added profile picture field

    // Relationship
    Account Accounts `gorm:"foreignKey:MetamaskAddress;references:MetamaskAddress"`
}

type Organization struct {
    MetamaskAddress string    `gorm:"primaryKey;size:42" json:"metamask_address"` // Same as Accounts
    AcadEmail       string    `gorm:"not null;unique" json:"acad_email"`
    OrgName         string    `gorm:"not null" json:"org_name"`
    OrgType         string    `gorm:"size:100" json:"org_type"` // "university", "college"
    OrgUrl          string    `gorm:"size:255" json:"org_url"`
    LogoIPFSHash    string    `gorm:"size:100" json:"logo_ipfs_hash"` // IPFS instead of bytes
    OrgDesc         string    `gorm:"type:text" json:"org_desc"`
     Country         string    `gorm:"size:100" json:"country"`
    State           string    `gorm:"size:100" json:"state"`
    City            string    `gorm:"size:100" json:"city"`
    Address         string    `gorm:"type:text" json:"address"`
    PostalCode      string    `gorm:"size:20" json:"postal_code"`
    PhoneNumber     string    `gorm:"size:20" json:"phone_number"`
    EstablishedYear int       `json:"established_year"`
    Accreditation   string    `gorm:"size:255" json:"accreditation"`
    IsVerified      bool      `gorm:"default:false" json:"is_verified"`
    
    // Statistics for dashboard
    TotalCredentials     int   `gorm:"default:0" json:"total_credentials"`
    ActiveCredentials    int   `gorm:"default:0" json:"active_credentials"`
    TotalStudents        int   `gorm:"default:0" json:"total_students"`
    
    CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// Relationships
	Account Accounts `gorm:"foreignKey:MetamaskAddress;references:MetamaskAddress"`
}

type CredentialTemplate struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    UniversityWallet string   `gorm:"not null;size:42;index" json:"university_wallet"`
    TemplateName    string    `gorm:"not null" json:"template_name"`
    DegreeType      string    `gorm:"not null" json:"degree_type"` // "Bachelor", "Master", "PhD", "Certificate"
    Major           string    `gorm:"not null" json:"major"`
    Department      string    `gorm:"size:255" json:"department"`
    CourseDuration  string    `gorm:"size:50" json:"course_duration"` // "4 years", "2 years"
    MinGPA          string    `gorm:"size:10" json:"min_gpa"`
    IsActive        bool      `gorm:"default:true" json:"is_active"`
    TemplateIPFS    string    `gorm:"size:100" json:"template_ipfs"` // IPFS hash for template design
    
    // Metadata for template
    RequiredFields  datatypes.JSON `gorm:"type:jsonb" json:"required_fields"`
    TemplateConfig  datatypes.JSON `gorm:"type:jsonb" json:"template_config"`
    
    CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relationship
    University Organization `gorm:"foreignKey:UniversityWallet;references:MetamaskAddress"`
}

type Credential struct {
    ID              string         `gorm:"primaryKey;size:255" json:"id"`
    StudentID       string         `gorm:"not null;size:100;index" json:"student_id"`
    StudentWallet   string         `gorm:"not null;size:42" json:"student_wallet"`
    UniversityID    string         `gorm:"not null;size:100;index" json:"university_id"`
    DegreeName      string         `gorm:"not null;size:255" json:"degree_name"`
    Major           string         `gorm:"size:255" json:"major"`
    GPA             string         `gorm:"size:10" json:"gpa"`
    GraduationDate  string         `gorm:"size:50" json:"graduation_date"`
    IPFSHash        string         `gorm:"unique;not null;size:100;index" json:"ipfs_hash"`
    TokenID         string         `gorm:"unique;size:100;index" json:"token_id"`
    ContractAddress string         `gorm:"size:42" json:"contract_address"`
    Status          string         `gorm:"default:Active;size:50;index" json:"status"`
    
	CertificateIPFS string         `gorm:"size:100" json:"certificate_ipfs"` // Actual certificate document
    TranscriptIPFS  string         `gorm:"size:100" json:"transcript_ipfs"`  // Academic transcript
    
    CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    Metadata        datatypes.JSON `gorm:"type:jsonb" json:"metadata"`

    Student     Users               `gorm:"foreignKey:StudentWallet;references:MetamaskAddress"`
    University  Organization        `gorm:"foreignKey:UniversityWallet;references:MetamaskAddress"`
    Template    CredentialTemplate  `gorm:"foreignKey:TemplateID;references:ID"`
}

type StudentAcademicInfo struct {
    // From Users table
    MetamaskAddress string `json:"metamask_address"`
    FirstName       string `json:"first_name"`
    LastName        string `json:"last_name"`
    Email           string `json:"email"`
    StudentID       string `json:"student_id"`
    
    // From Credential table
    Major           string `json:"major"`
    GPA             string `json:"gpa"`
    GraduationDate  string `json:"graduation_date"`
    Status          string `json:"status"`
    
    // From CredentialTemplate table
    DegreeType      string `json:"degree_type"`
    Department      string `json:"department"`
    CourseDuration  string `json:"course_duration"`
    
    // Derived fields
    TotalCredentials int   `json:"total_credentials"`
    EnrollmentStatus string `json:"enrollment_status"` // "graduated", "enrolled"
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

type Employer struct {
    MetamaskAddress string    `gorm:"primaryKey;size:42" json:"metamask_address"`
    CompanyName     string    `gorm:"not null" json:"company_name"`
    CompanyEmail    string    `gorm:"unique;not null" json:"company_email"`
    CompanyWebsite  string    `json:"company_website"`
    Industry        string    `gorm:"size:100" json:"industry"`
    LogoIPFSHash    string    `gorm:"size:100" json:"logo_ipfs_hash"`
    IsVerified      bool      `gorm:"default:false" json:"is_verified"`
    CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    
    Account Accounts `gorm:"foreignKey:MetamaskAddress;references:MetamaskAddress"`
}

type VerificationRequest struct {
    ID               uint      `gorm:"primaryKey" json:"id"`
    EmployerWallet   string    `gorm:"not null;size:42" json:"employer_wallet"`
    StudentWallet    string    `gorm:"not null;size:42" json:"student_wallet"`
    CredentialID     string    `gorm:"not null" json:"credential_id"`
    RequestMessage   string    `gorm:"type:text" json:"request_message"`
    Status           string    `gorm:"default:pending" json:"status"` // "pending", "approved", "denied"
    RequestedAt      time.Time `gorm:"autoCreateTime" json:"requested_at"`
    RespondedAt      *time.Time `json:"responded_at"`

    
    Employer       Employer   `gorm:"foreignKey:EmployerWallet;references:MetamaskAddress"`
    Student        Users      `gorm:"foreignKey:StudentWallet;references:MetamaskAddress"`
    Credential     Credential `gorm:"foreignKey:CredentialID;references:ID"`
}

type UniversityAdmin struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    MetamaskAddress string    `gorm:"unique;not null;size:42" json:"metamask_address"`
    UniversityWallet string   `gorm:"not null;size:42;index" json:"university_wallet"`
    Email           string    `gorm:"not null;unique" json:"email"`
    FirstName       string    `gorm:"not null" json:"first_name"`
    LastName        string    `gorm:"not null" json:"last_name"`
    Role            string    `gorm:"not null" json:"role"` 
    Department      string    `gorm:"size:255" json:"department"`
    Permissions     datatypes.JSON `gorm:"type:jsonb" json:"permissions"`
    IsActive        bool      `gorm:"default:true" json:"is_active"`
    
    CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

  
    University Organization `gorm:"foreignKey:UniversityWallet;references:MetamaskAddress"`
}

type UniversityDashboardStats struct {
    TotalStudents       int64 `json:"total_students"`
    ActiveCredentials   int64 `json:"active_credentials"`
    TotalCredentials    int64 `json:"total_credentials"`
    GraduatedStudents   int64 `json:"graduated_students"`
    PendingCredentials  int64 `json:"pending_credentials"`
    
    // By Department
    DepartmentStats []DepartmentStat `json:"department_stats"`
    
    // By Degree Type
    DegreeTypeStats []DegreeTypeStat `json:"degree_type_stats"`
}

type DepartmentStat struct {
    Department      string `json:"department"`
    StudentCount    int    `json:"student_count"`
    CredentialCount int    `json:"credential_count"`
}

type DegreeTypeStat struct {
    DegreeType      string `json:"degree_type"`
    StudentCount    int    `json:"student_count"`
}

func GetUniversityDashboardStats(universityWallet string) (*UniversityDashboardStats, error) {
    stats := &UniversityDashboardStats{}
    
    // Total students
    DB.Table("credentials").
        Where("university_wallet = ?", universityWallet).
        Distinct("student_wallet").
        Count(&stats.TotalStudents)
    
    // Active credentials
    DB.Table("credentials").
        Where("university_wallet = ? AND status = ?", universityWallet, "Active").
        Count(&stats.ActiveCredentials)
    
    // Total credentials
    DB.Table("credentials").
        Where("university_wallet = ?", universityWallet).
        Count(&stats.TotalCredentials)
    
    // Graduated students (have graduation date in past)
    DB.Table("credentials").
        Where("university_wallet = ? AND graduation_date <= NOW() AND status = ?", universityWallet, "Active").
        Distinct("student_wallet").
        Count(&stats.GraduatedStudents)
    
    // Department stats
    DB.Table("credentials c").
        Select("ct.department, COUNT(DISTINCT c.student_wallet) as student_count, COUNT(c.id) as credential_count").
        Joins("JOIN credential_templates ct ON c.template_id = ct.id").
        Where("c.university_wallet = ?", universityWallet).
        Group("ct.department").
        Find(&stats.DepartmentStats)
    
    // Degree type stats
    DB.Table("credentials c").
        Select("ct.degree_type, COUNT(DISTINCT c.student_wallet) as student_count").
        Joins("JOIN credential_templates ct ON c.template_id = ct.id").
        Where("c.university_wallet = ?", universityWallet).
        Group("ct.degree_type").
        Find(&stats.DegreeTypeStats)
    
    return stats, nil
}

type UserDashboardProfile struct {
    // From Users
    MetamaskAddress     string    `json:"metamask_address"`
    Email               string    `json:"email"`
    FirstName           string    `json:"first_name"`
    LastName            string    `json:"last_name"`
    StudentID           string    `json:"student_id"`
    ProfilePicture      string    `json:"profile_picture"`
    Bio                 string    `json:"bio"`
    PhoneNumber         string    `json:"phone_number"`
    
    // From Accounts
    AccountType         string    `json:"account_type"`
    Verified            bool      `json:"verified"`
    TotalCredentials    int       `json:"total_credentials"`
    LastLoginAt         *time.Time `json:"last_login_at"`
    
    // Derived from Credentials
    UniversitiesAttended []string  `json:"universities_attended"`
    LatestDegree        string    `json:"latest_degree"`
    LatestGPA           string    `json:"latest_gpa"`
    CredentialCount     int64     `json:"credential_count"`
}

func GetUserDashboardProfile(userWallet string) (*UserDashboardProfile, error) {
    profile := &UserDashboardProfile{}
    
    // Get user basic info
    var user Users
    if err := DB.Where("metamask_address = ?", userWallet).First(&user).Error; err != nil {
        return nil, err
    }
    
    // Get account info
    var account Accounts
    DB.Where("metamask_address = ?", userWallet).First(&account)
    
    // Map data
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
    
    // Get credential-derived data
    var credentials []Credential
    DB.Where("student_wallet = ?", userWallet).Find(&credentials)
    
    profile.CredentialCount = int64(len(credentials))
    
    // Get universities attended
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
    
    // Get latest degree info
    if len(credentials) > 0 {
        latest := credentials[len(credentials)-1] // Most recent
        profile.LatestDegree = latest.DegreeName
        profile.LatestGPA = latest.GPA
    }
    
    return profile, nil
}