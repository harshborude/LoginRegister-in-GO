package db

import (
	"backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func seedAdmin() {

	var admin models.User

	result := DB.Where("email = ?", "admin@auction.com").First(&admin)

	if result.Error == gorm.ErrRecordNotFound {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("error occurred during admin password hashing: %v", err)
			return
		}

		adminUser := models.User{
			Username:     "admin",
			Email:        "admin@auction.com",
			PasswordHash: string(hashedPassword),
			Role:         "ADMIN",
		}

		if err := DB.Create(&adminUser).Error; err != nil {
			log.Printf("error occurred during admin creation: %v", err)
			return
		}

		log.Println("Default admin created")
	}
}

func ConnectDatabase() {

	dsn := "host=localhost user=postgres password=harsh123 dbname=auction port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = database

	DB.AutoMigrate(&models.User{})

	seedAdmin()

	log.Println("Database connected successfully")
}