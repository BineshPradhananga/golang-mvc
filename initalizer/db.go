package initalizer

import (
	"os"

	"github.com/binesh/gomvc/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil // Return the existing connection if it's already initialized
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return DB, nil
}

func SyncDb(DB *gorm.DB) {
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})
}
