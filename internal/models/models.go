package models

import (
	"time"
)


type Accounts struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	MetamaskAddress string    `json:"metamask_address"` // Ethereum address from Metamask wallet
	AccountType     string    `json:"account_type"`     // "individual" or "organization"
	CreatedAt       time.Time `json:"created_at"`
	// Password        string    `json:"-"` // we dont this for now, if we're verifying shit by jwt
	Verified        bool      `default:"false" json:"verified"`
}

type Users struct {
	MetamaskAddress string `json:"metamask_address"` // Ethereum address from Metamask wallet
	Email           string `gorm:"notNull" json:"email"`
	FirstName       string `gorm:"notNull" json:"first_name"`
	LastName        string `gorm:"notNull" json:"last_name"`
}

type Organization struct {
	MetamaskAddress string    `json:"metamask_address"`
	AcadEmail       string    `gorm:"notNull" json:"acad_email"`
	OrgName         string    `gorm:"notNull" json:"org_name"`
	OrgType         string    `gorm:"notNull" json:"org_type"` // "college", "school", "institute", etc.
}