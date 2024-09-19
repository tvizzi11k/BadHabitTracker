package models_

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Habits   []Habit
}
