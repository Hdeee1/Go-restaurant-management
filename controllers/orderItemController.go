package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
)

// GetOrderItems godoc
// @Summary Get all order items
// @Description Retrieve a paginated list of all order items
// @Tags OrderItems
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orderitems [get]
func GetOrderItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderItems []models.OrderItem

		result := database.DB.Scopes(helpers.Paginate(ctx)).Find(&orderItems)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"order_items": orderItems,
			"page":        ctx.DefaultQuery("page", "1"),
			"limit":       ctx.DefaultQuery("limit", "10"),
		})
	}
}

// GetOrderItem godoc
// @Summary Get order item by ID
// @Description Retrieve a specific order item by order_item_id
// @Tags OrderItems
// @Accept json
// @Produce json
// @Param order_item_id path string true "Order Item ID"
// @Security BearerAuth
// @Success 200 {object} models.OrderItem
// @Failure 404 {object} map[string]interface{}
// @Router /orderitems/{order_item_id} [get]
func GetOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		order_item_id := ctx.Param("order_item_id")

		var orderItem models.OrderItem

		if err := database.DB.Where("order_item_id = ?", order_item_id).First(&orderItem).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "order_item_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, orderItem)
	}
}

// UpdateOrderItem godoc
// @Summary Update an order item
// @Description Update an existing order item by order_item_id
// @Tags OrderItems
// @Accept json
// @Produce json
// @Param order_item_id path string true "Order Item ID"
// @Param order_item body models.OrderItem true "Order item object"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orderitems/{order_item_id} [put]
func UpdateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
			"message":       "order item updated",
			"order_item_id": orderItem.Order_item_id,
		})
	}
}
