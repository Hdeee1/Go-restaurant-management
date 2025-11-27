package controllers

import (
	"net/http"
	"time"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderRequest struct {
	Table_id    string             `json:"table_id" validate:"required"`
	Order_items []models.OrderItem `json:"order_items" validate:"required"`
}

// GetOrders godoc
// @Summary Get all orders
// @Description Retrieve a paginated list of all orders with their order items
// @Tags Orders
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders [get]
func GetOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orders []models.Order

		result := database.DB.Scopes(helpers.Paginate(ctx)).Preload("OrderItems").Find(&orders)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"orders": orders,
			"page":   ctx.DefaultQuery("page", "1"),
			"limit":  ctx.DefaultQuery("limit", "10"),
		})
	}
}

// GetOrder godoc
// @Summary Get order by ID
// @Description Retrieve a specific order by order_id with its order items
// @Tags Orders
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Security BearerAuth
// @Success 200 {object} models.Order
// @Failure 404 {object} map[string]interface{}
// @Router /orders/{order_id} [get]
func GetOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		order_id := ctx.Param("order_id")

		var order models.Order

		if err := database.DB.Preload("OrderItems").Where("order_id = ?", order_id).First(&order).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "order_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, order)
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with order items. Automatically validates food items and calculates prices.
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body controllers.OrderRequest true "Order with items"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders [post]
func CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var order models.Order
		var req OrderRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := helpers.Validate.Struct(req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tx := database.DB.Begin()

		order.Order_id = uuid.New().String()
		order.Order_date = time.Now()
		order.Table_id = &req.Table_id

		if err := tx.Create(&order).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, item := range req.Order_items {
			var food models.Food
			if err := database.DB.Where("food_id = ?", item.Food_id).First(&food).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusNotFound, gin.H{"error": "food_id not found"})
				return
			}

			item.Unit_price = food.Price
			item.Order_item_id = uuid.New().String()
			item.Order_id = order.Order_id

			if item.Food_id == nil || item.Quantity == nil {
				tx.Rollback()
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incomplete data item"})
				return
			}

			if err := tx.Create(&item).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order item"})
				return
			}
		}
		tx.Commit()

		ctx.JSON(http.StatusCreated, gin.H{
			"message":  "order created",
			"order_id": order.Order_id,
		})
	}
}

// CreateOrderItem godoc
// @Summary Create a new order item
// @Description Add a new item to an existing order
// @Tags Orders
// @Accept json
// @Produce json
// @Param order_item body models.OrderItem true "Order item object"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orderitems [post]
func CreateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
			"message":       "order item created",
			"order_item_id": orderItem.Order_item_id,
		})
	}
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update an existing order by order_id. Cannot update if the order is already paid.
// @Tags Orders
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Param order body models.Order true "Order object"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders/{order_id} [put]
func UpdateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		order_id := ctx.Param("order_id")

		var order models.Order

		if err := database.DB.Where("order_id = ?", order_id).First(&order).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "order_id not found"})
			return
		}

		var invoice models.Invoice

		err := database.DB.Where("order_id = ?", order_id).First(&invoice).Error
		if err == nil {
			if invoice.Payment_status != nil && *invoice.Payment_status == "PAID" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "order already paid"})
				return
			}
		}

		if err := helpers.Validate.Struct(order); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
			"message":  "order updated",
			"order_id": order.Order_id,
		})
	}
}
