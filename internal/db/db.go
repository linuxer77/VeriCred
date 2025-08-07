package db

import (
	"fmt"
	"log"
	_ "os"
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
	err = DB.AutoMigrate(&models.Users{}, &models.Organization{}, &models.Accounts{})
	if err != nil {
		log.Fatal("AutoMigration failed: ", err)
	}
}
