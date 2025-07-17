package models

import (
	"time"
)


type Accounts struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	MetamaskAddress string    `json:"metamask_address"` // Ethereum address from Metamask wallet
	AccountType     string    `json:"account_type"`     // "individual" or "organization"
	CreatedAt       time.Time `json:"created_at"`
	Password        string    `json:"-"` // Hide from JSON
	Verified        bool      `json:"verified"`
}

type Users struct {
	MetamaskAddress string `json:"metamask_address"` // Ethereum address from Metamask wallet
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
}

type Organization struct {
	MetamaskAddress string    `json:"metamask_address"`
	AcadEmail       string    `json:"acad_email"` 
	OrgName         string    `json:"org_name"` 
	OrgType         string    `json:"org_type"` // "college", "school", "institute", etc.
}