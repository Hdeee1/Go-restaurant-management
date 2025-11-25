package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Role			*string		`json:"role" validate:"required"`
	First_name		*string		`json:"first_name" validate:"required,min=2,max=100"`
	Last_name		*string		`json:"last_name" validate:"required,min=2,max=100"`
	Password		*string		`json:"password" validate:"required,min=8"`
	Email			*string		`json:"email" validate:"email,required"`
	Avatar			*string		`json:"avatar"`
	Phone			*string		`json:"phone" validate:"required"`
	Token			*string		`json:"token"`
	Refresh_Token	*string		`json:"refresh_token"`
	User_id			string		`json:"user_id"`
}