package controllers

import (
	"net/http"
	
	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetOrders() gin.HandlerFunc {
	return  func(ctx *gin.Context) {

	}
}

func GetOrder() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		
	}
}

func CreateOrder() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		var order models.Order

		order.Order_id = uuid.New().String()

		if err := ctx.BindJSON(&order); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Create(&order).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "order created",
			"order_id": order.Order_id,
		})
	}
}

func UpdateOrder() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		
	}
}