package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name       *string  `json:"name" validate:"required,min=2,max=100"`
	Price      *float64 `json:"price" validate:"required"`
	Food_image *string  `json:"food_image" validate:"required"`
	Food_id    string   `json:"food_id"`
	Menu_id    *string  `json:"menu_id" validate:"required"`
}
