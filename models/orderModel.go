package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Order_date		time.Time	`json:"order_date" validate:"required"`
	Order_id		string		`json:"order_id"`
	Table_id		*string		`json:"table_id"`
	OrderItems		[]OrderItem	`gorm:"foreignKey:Order_id;references:Order_id" json:"order_items"`
}