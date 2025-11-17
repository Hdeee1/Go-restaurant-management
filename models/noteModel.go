package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Text		string	`json:"text"`
	Title		string	`json:"title"`
	Note_id		string	`json:"note_id"`
}