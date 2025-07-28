package models

import (
	"time"
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

// type Credential struct {
// 	ID             uint         `gorm:"primaryKey"`
// 	Name           string       `json:"name"`
// 	Type           string       `json:"type"`
// 	Description    string       `json:"description"`
// 	OrganizationID uint
// 	Org            Organization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }
