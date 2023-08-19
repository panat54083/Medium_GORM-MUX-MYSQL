package config

import (
	"example/backend/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB(dsn string) {
	var err error
	// Open a connection to the MySQL database using the provided DSN.
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d
	// Migrate the schema
	db.AutoMigrate(&models.User{}) // User

	fmt.Println("Connected to MySQL!")
}

// GetDB returns the initialized database connection.
func GetDB() *gorm.DB {
	return db
}
