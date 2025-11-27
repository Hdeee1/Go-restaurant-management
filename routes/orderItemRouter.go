package routes

import (
	"github.com/Hdeee1/go-restaurant-management/controllers"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/orderItems", middleware.Authentication(), controllers.CreateOrderItem())
	incomingRoutes.GET("/orderItems", middleware.Authentication(), controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", middleware.Authentication(), controllers.GetOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", middleware.Authentication(), controllers.UpdateOrderItem())
}
