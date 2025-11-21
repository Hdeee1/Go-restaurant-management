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
		var foods []models.Food

		err := database.DB.Find(&foods).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return 
		}

		ctx.JSON(http.StatusOK, gin.H{
			"foods": foods,
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		foodID := ctx.Param("food_id")

		var food models.Food

		if err := database.DB.Where("food_id = ?", foodID).Find(&food).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "food_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"food_id": foodID})
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
