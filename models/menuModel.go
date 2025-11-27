package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name       string     `json:"name" validate:"required"`
	Category   string     `json:"category" validate:"required"`
	Start_date *time.Time `json:"start_date"`
	End_date   *time.Time `json:"end_date"`
	Menu_id    string     `json:"menu_id" validate:"required"`
}
