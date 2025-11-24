package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hdeee1/go-restaurant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrder())
	incomingRoutes.PATCH("/orders/:order_id", controllers.UpdateOrder())
}