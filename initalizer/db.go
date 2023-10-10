package initalizer

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/binesh/gomvc/models"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:P@ssw0rd@tcp(DB_HOST=hmvcdb:33067)/godb"
	dsn := os.Getenv("DB_URL")
	// DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database", err)
	}
	fmt.Println(err)
	// fmt.Println(DB)

}

func SyncDb() {
	// DB.AutoMigrate(&models.Post{})
}
