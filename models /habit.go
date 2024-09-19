package models_

import "github.com/jinzhu/gorm"

type Habit struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Goal        int    `json:"goal"`
	UserID      uint   `json:"user_id"`
	Trackings   []Tracking
}
