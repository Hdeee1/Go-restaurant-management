package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hdeee1/go-restaurant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/table", controllers.CreateTable())
	incomingRoutes.GET("/table", controllers.GetTables())
	incomingRoutes.GET("/table/:table_id", controllers.GetTable())
	incomingRoutes.PATCH("/table/:table_id", controllers.UpdateTable())
}