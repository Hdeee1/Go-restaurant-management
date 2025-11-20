package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}

func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func AddFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var food models.Food
		
		if err := ctx.BindJSON(&food); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		food.Food_id = uuid.New().String()

		var menu models.Menu
		err := database.DB.Where("menu_id = ?", *food.Menu_id).First(&menu).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "menu_id not found"})
			return
		}

		if err := database.DB.Create(&food).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Food created",
			"food_id": food.Food_id,
		})		
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}
