package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	models_ "habit-tracker/models "
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate(&models_.User{}, &models_.Habit{}, &models_.Tracking{})
	DB = db
}
