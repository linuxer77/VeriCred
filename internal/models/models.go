package models

import (
	"time"
	"gorm.io/datatypes"
)

type Accounts struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	MetamaskAddress string    `json:"metamask_address"` // eth adrs from metamask wallet
	AccountType     string    `json:"account_type"`     // "individual" or "organization"
	// Password        string    `json:"-"` // we dont this for now, if we're verifying shit by jwt
	Verified        bool      `default:"false" json:"verified"`
	Credentials	 	int	 	  `default:"0" json:"credentails"` // Total shit minted
	CreatedAt       time.Time `json:"created_at"`
}

type Users struct {
	MetamaskAddress string `json:"metamask_address"` 
	Email           string `gorm:"notNull" json:"email"`
	FirstName       string `gorm:"notNull" json:"first_name"`
	LastName        string `gorm:"notNull" json:"last_name"`
	Password 		string `gorm:"notNull" json:"password"`
}

type Organization struct {
	ID              uint           `gorm:"primaryKey"`
	MetamaskAddress string         `json:"metamask_address"`
	AcadEmail       string         `gorm:"not null" json:"acad_email"`
	OrgName         string         `gorm:"not null" json:"org_name"`
	OrgType         *string        `json:"org_type"`
	OrgUrl          *string        `json:"org_url"`
	OrgImg          []byte         `json:"org_img"`
	// Credentials     []Credential   `json:"certificate"`
	OrgDesc         *string        `json:"org_desc"`
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
    CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    Metadata        datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
}

type University struct {
    ID            string    `gorm:"primaryKey;size:100" json:"id"`
    Name          string    `gorm:"not null;size:255" json:"name"`
    WalletAddress string    `gorm:"not null;size:42" json:"wallet_address"`
    LogoIPFSHash  string    `gorm:"size:100" json:"logo_ipfs_hash"`
    IsVerified    bool      `gorm:"default:false" json:"is_verified"`
    CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
// type Credential struct {
// 	ID             uint         `gorm:"primaryKey"`
// 	Name           string       `json:"name"`
// 	Type           string       `json:"type"`
// 	Description    string       `json:"description"`
// 	OrganizationID uint
// 	Org            Organization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }
