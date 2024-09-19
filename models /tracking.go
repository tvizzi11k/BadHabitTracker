package models_

import "github.com/jinzhu/gorm"

type Tracking struct {
	gorm.Model
	HabitID uint   `json:"habit_id"`
	Note    string `json:"note"`
}
