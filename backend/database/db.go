package database

import (
	"log"

	"github.com/pump-p/solidithai-assignment-2/backend/config"
	"github.com/pump-p/solidithai-assignment-2/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	cfg := config.LoadConfig()
	dsn := config.GetDSN(cfg)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Database connected successfully.")

	// Migrate the User model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
