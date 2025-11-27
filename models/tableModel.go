package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Number_of_guest *int   `json:"number_of_guest" validate:"required"`
	Table_number    *int   `json:"table_number" validate:"required"`
	Table_id        string `json:"table_id"`
}
