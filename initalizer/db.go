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
	// dsn := "root:P@ssw0rd@tcp(DB_HOST=hmvcdb:33067)/godb"
	dsn := os.Getenv("DB_URL")
	// dsn := "root:P@ssw0rd@tcp(localhost:33067)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database", err)
	}
	// fmt.Println(DB)

}

func SyncDb() {
	DB.AutoMigrate(&models.Post{})
}
