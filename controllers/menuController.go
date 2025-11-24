package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menus []models.Menu

		err := database.DB.Find(&menus).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return 
		}

		ctx.JSON(http.StatusOK, gin.H{
			"menus": menus,
		})
	}
}

func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		menu_id := ctx.Param("menu_id")

		var menu models.Menu

		if err := database.DB.Where("menu_id = ?", menu_id).First(&menu).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "menu_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menu models.Menu

		menu.Menu_id = uuid.New().String()

		if err := ctx.BindJSON(&menu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		if menu.Start_date != nil && menu.End_date != nil {
			if menu.End_date.Before(*menu.Start_date) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "end_date must be after start_date"})
				return
			}
		}

		if err := database.DB.Create(&menu).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "menu created",
			"menu_id": menu.Menu_id,
		})
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}