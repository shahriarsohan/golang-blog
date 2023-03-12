package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shahriarsohan/go-blog-practise/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error connecting to database")
	} else {
		log.Println("Connected successfully to DB")
	}

	DB = db
	DB.AutoMigrate(
		&models.User{},
	)
}
