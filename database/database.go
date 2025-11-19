package database

import (
	"log"

	"github.com/Hdeee1/go-restaurant-management/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:dbmysql@tcp(localhost:3306)/go_restaurant_management?parseTime=true"
	
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.User{})
}
