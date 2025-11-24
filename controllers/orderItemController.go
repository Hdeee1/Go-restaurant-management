package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetOrderItems() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		var orderItems []models.OrderItem

		if err := database.DB.Find(&orderItems).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"order_items": orderItems,
		})
	}
}

func GetOrderItem() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		order_item_id := ctx.Param("order_item_id")

		var orderItem models.OrderItem

		if err := database.DB.Where("order_item_id = ?", order_item_id).First(&orderItem).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "order_item_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, orderItem)
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		var orderItem models.OrderItem
		
		if err := ctx.BindJSON(&orderItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := helpers.Validate.Struct(orderItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		orderItem.Order_item_id = uuid.New().String()

		if err := database.DB.Create(&orderItem).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "order item created",
			"order_item_id": orderItem.Order_item_id,
		})
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		order_item_id := ctx.Param("order_item_id")

		var orderItem models.OrderItem

		if err := database.DB.Where("order_item_id = ?", order_item_id).First(&orderItem).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "order_item_id not found"})
			return
		}

		var updateData models.OrderItem
		if err := ctx.BindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Model(&orderItem).Updates(updateData).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "order item updated",
			"order_item_id": orderItem.Order_item_id,
		})
	}
}