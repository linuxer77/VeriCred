package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"vericred/internal/db"
	"vericred/pkg/ipfs"
	"vericred/pkg/nft"
	"vericred/pkg/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

type CredentialRequest struct {
	UniversityID   string `json:"university_id"`
	StudentID      string `json:"student_id"`
	StudentName    string `json:"student_name"`
	StudentWallet  string `json:"student_wallet"`
	DegreeName     string `json:"degree_name"`
	Major          string `json:"major"`
	CredentialType string `json:"credential_type"`
	GPA            string `json:"gpa"`
	GraduationDate string `json:"graduation_date"`
	Description    string `json:"description"`
}

type NFTMetadata struct {
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	Image          string                 `json:"image"`
	ExternalURL    string                 `json:"external_url"`
	Attributes     []Attribute            `json:"attributes"`
	BlockchainData map[string]interface{} `json:"blockchain_data"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

func MintCredential(w http.ResponseWriter, r *http.Request) {
	// 1. Parse university request
	var req CredentialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// 2. Validate university permissions
	university, err := db.GetUniversity(req.UniversityID)
	if err != nil {
		http.Error(w, "University not found", http.StatusNotFound)
		return
	}
	if !university.IsVerified {
		http.Error(w, "University not authorized", http.StatusForbidden)
		return
	}

	// 3. Create comprehensive NFT metadata
	credentialID := generateCredentialID(university.Name, req.Major, req.StudentID)
	currentTime := time.Now().Format("2006-01-02")

	metadata := NFTMetadata{
		Name:        req.DegreeName,
		Description: req.Description,
		Image:       fmt.Sprintf("ipfs://%s/university_logo.png", university.LogoIPFSHash),
		ExternalURL: fmt.Sprintf("ipfs://%s/credential_details.pdf", ""), // Will be updated after PDF upload
		Attributes: []Attribute{
			{TraitType: "Credential Type", Value: req.CredentialType},
			{TraitType: "Issuing Institution", Value: university.Name},
			{TraitType: "Issuer Wallet", Value: university.WalletAddress},
			{TraitType: "Recipient Name", Value: req.StudentName},
			{TraitType: "Recipient Wallet", Value: req.StudentWallet},
			{TraitType: "Issue Date", Value: currentTime},
			{TraitType: "Graduation Date", Value: req.GraduationDate},
			{TraitType: "Major", Value: req.Major},
			{TraitType: "GPA", Value: req.GPA},
			{TraitType: "Credential ID", Value: credentialID},
			{TraitType: "Status", Value: "Active"},
			{TraitType: "Verification URL", Value: fmt.Sprintf("https://vericred.com/verify?credentialId=%s", credentialID)},
		},
		BlockchainData: map[string]interface{}{
			"contractAddress": "",    // Will be filled after minting
			"tokenId":         "",    // Will be filled after minting
			"chainId":         "137", // Polygon mainnet
		},
	}

	// 4. Upload JSON metadata to IPFS
	jsonData, err := json.Marshal(metadata)
	if err != nil {
		http.Error(w, "Error creating metadata", http.StatusInternalServerError)
		return
	}

	ipfsHash, err := ipfs.Add(jsonData)
	if err != nil {
		http.Error(w, "Error uploading to IPFS", http.StatusInternalServerError)
		return
	}

	// 5. Mint NFT with metadata hash
	tokenID, contractAddress, err := nft.Mint(req.StudentWallet, ipfsHash)
	if err != nil {
		http.Error(w, "Error minting NFT", http.StatusInternalServerError)
		return
	}

	// 6. Store credential record in database
	// Update your StoreCredential function
	err = db.StoreCredential(db.GetCredential{
		ID:              credentialID,
		StudentID:       req.StudentID,
		StudentWallet:   req.StudentWallet,
		UniversityID:    req.UniversityID,
		DegreeName:      req.DegreeName,
		Major:           req.Major,
		GPA:             req.GPA,
		GraduationDate:  req.GraduationDate,
		IPFSHash:        ipfsHash,
		TokenID:         tokenID,
		ContractAddress: contractAddress,
		Status:          "Active",
		CreatedAt:       time.Now(),
		// Store the full JSON metadata in PostgreSQL JSONB
		Metadata: postgres.Jsonb{RawMessage: jsonData},
	})

	// Also cache in Redis for fast lookups
	credentialCache := CredentialCache{
		IPFSHash:     ipfsHash,
		TokenID:      tokenID,
		Status:       "Active",
		LastAccessed: time.Now(),
	}
	redisClient.HSet(ctx, fmt.Sprintf("credential:%s", credentialID), credentialCache)
	redisClient.HSet(ctx, fmt.Sprintf("ipfs:%s", ipfsHash), credentialID)
	// 7. Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":          true,
		"credential_id":    credentialID,
		"ipfs_hash":        ipfsHash,
		"token_id":         tokenID,
		"contract_address": contractAddress,
		"verification_url": fmt.Sprintf("https://vericred.com/verify?credentialId=%s", credentialID),
	})
}

func generateCredentialID(universityName, major, studentID string) string {
	// Generate unique credential ID
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s-%s-%s-%d",
		abbreviateUniversity(universityName),
		abbreviateMajor(major),
		studentID,
		timestamp)
}

func abbreviateUniversity(name string) string {
	// Simple abbreviation logic - you can make this more sophisticated
	if len(name) > 10 {
		return name[:3] + "U"
	}
	return name
}

func abbreviateMajor(major string) string {
	// Simple abbreviation logic
	if len(major) > 10 {
		return major[:2] + "S"
	}
	return major
}
