package initalizer

import (
	"fmt"
	"os"

	"github.com/binesh/gomvc/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database", err)
	}
	// fmt.Println(DB)

}

func SyncDb() {
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})
}
