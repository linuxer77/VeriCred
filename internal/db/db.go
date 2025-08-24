package db

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"vericred/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ensureSSLMode(dsn string) string {
	if strings.Contains(dsn, "postgres://") || strings.Contains(dsn, "postgresql://") {
		// ensure sslmode=require for hosted providers like Render
		if !strings.Contains(dsn, "sslmode=") {
			sep := "?"
			if strings.Contains(dsn, "?") {
				sep = "&"
			}
			dsn = dsn + sep + "sslmode=require"
		}
	}
	return dsn
}

func Init() {
	// Prefer DB_URL if provided (e.g., Render external URL)
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		// Fallback to local development DSN
		dsn = "host=localhost user=postgres password=post4321 dbname=vericred port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	} else {
		dsn = ensureSSLMode(dsn)
	}

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
	if err = DB.AutoMigrate(&models.Accounts{}); err != nil {
		log.Fatal("AutoMigration failed for Accounts: ", err)
	}
	if err = DB.AutoMigrate(&models.Organization{}); err != nil {
		log.Fatal("AutoMigration failed for Organization: ", err)
	}
	if err = DB.AutoMigrate(&models.Users{}); err != nil {
		log.Fatal("AutoMigration failed for Users: ", err)
	}
	if err = DB.AutoMigrate(&models.PendingRequest{}); err != nil {
		log.Fatal("AutoMigration failed for PendingRequest: ", err)
	}
	if err = DB.AutoMigrate(&models.Credential{}); err != nil {
		log.Fatal("AutoMigration failed for Credential: ", err)
	}
}
