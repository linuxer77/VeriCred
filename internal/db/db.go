package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"vericred/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=post4321 dbname=vericred port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to db failed:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get db from GORM: ", err)
	}
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("(SUCCESS): connected to database successfully ")
	
	// Initialize models DB connection
	models.InitDB(DB)
	
	// Check if we should use SQL schema or GORM AutoMigrate
	useSchema := os.Getenv("USE_SQL_SCHEMA")
	if useSchema == "true" {
		fmt.Println("Using SQL schema file for database setup")
		fmt.Println("Please run: psql -d vericred -f sql/schema/004_complete_schema.sql")
		return
	}
	
	// Default: Use GORM AutoMigrate for development
	err = DB.AutoMigrate(&models.Accounts{})
	if err != nil {
		log.Fatal("AutoMigration failed for Accounts: ", err)
	}
	
	err = DB.AutoMigrate(&models.Organization{})
	if err != nil {
		log.Fatal("AutoMigration failed for Organization: ", err)
	}
	
	err = DB.AutoMigrate(&models.Users{})
	if err != nil {
		log.Fatal("AutoMigration failed for Users: ", err)
	}
	
	// Pending requests table (required by pendingrequest handlers)
	err = DB.AutoMigrate(&models.PendingRequest{})
	if err != nil {
		log.Fatal("AutoMigration failed for PendingRequest: ", err)
	}
	
	// Credentials (uses FKs to users/orgs)
	err = DB.AutoMigrate(&models.Credential{})
	if err != nil {
		log.Fatal("AutoMigration failed for Credential: ", err)
	}

	err = DB.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatal("AutoMigration failed for Credential: ", err)
	}
	
	// err = DB.AutoMigrate(&models.VerificationRequest{})
	// if err != nil {
	// 	log.Fatal("AutoMigration failed for VerificationRequest: ", err)
	// }
}
