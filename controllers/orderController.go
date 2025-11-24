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
		var orders []models.Order

		if err := database.DB.Find(&orders).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"orders": orders,
		})
	}
}

func GetOrder() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		order_id := ctx.Param("order_id")

		var order models.Order

		if err := database.DB.Where("order_id = ?", order_id).First(&order).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "order_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, order)
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
		order_id := ctx.Param("order_id")

		var order models.Order

		if err := database.DB.Where("order_id = ?", order_id).First(&order).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "order_id not found"})
			return
		}

		var updateData models.Order
		if err := ctx.BindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Model(&order).Updates(updateData).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "order updated",
			"order_id": order.Order_id,
		})
	}
}