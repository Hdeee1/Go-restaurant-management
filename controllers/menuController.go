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


	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}