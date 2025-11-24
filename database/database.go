package database

import (
	"log"
	"os"

	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	godotenv.Load()
	
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"
	
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(
		&models.User{},
		&models.Food{},
		&models.Invoice{},
		&models.Menu{},
		&models.Note{},
		&models.Order{},
		&models.OrderItem{},
		&models.Table{},
	)
}
