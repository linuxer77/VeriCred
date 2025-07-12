package models

import (
	"time"
)


type User struct {
	MetamaskAddress string    `json:"metamask_address"` // Ethereum address from Metamask wallet
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"created_at"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Password        string    `json:"password"`
}

type Organization struct {
	MetamaskAddress string    `json:"metamask_address"`
	AcadEmail       string    `json:"acad_email"` 
	CreatedAt       time.Time `json:"created_at"`
	OrgName         string    `json:"org_name"` 
	Password        string    `json:"password"` 
}