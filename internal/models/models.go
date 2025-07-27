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
	MetamaskAddress string			      `json:"metamask_address"`
	AcadEmail       string    			  `gorm:"notNull" json:"acad_email"`
	OrgName         string    			  `gorm:"notNull" json:"org_name"`
	OrgType         string    			  `gorm:"Null" json:"org_type"` 
	OrgUrl 			string 	  			  `gorm:"Null" json:"org_url"`
	OrgImg 			[]byte	  			  `gorm:"Null" json:"org_img"`
	Credentials		[]Credential	 	  `gorm:"Null" json:"certificate"`
	OrgDesc			string				  `gorm:"Null" json:"org_desc"`
}

type Credential struct {
	Name 			string
	Type 			string
	Description		string
}
