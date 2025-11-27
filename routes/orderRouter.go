package routes

import (
	"github.com/Hdeee1/go-restaurant-management/controllers"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/orders", middleware.Authentication(), controllers.CreateOrder())
	incomingRoutes.GET("/orders", middleware.Authentication(), controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", middleware.Authentication(), controllers.GetOrder())
	incomingRoutes.PATCH("/orders/:order_id", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.UpdateOrder())
}
