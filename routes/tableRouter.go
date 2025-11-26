package routes

import (
	"github.com/Hdeee1/go-restaurant-management/controllers"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/table", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.CreateTable())
	incomingRoutes.GET("/table", middleware.Authentication(), controllers.GetTables())
	incomingRoutes.GET("/table/:table_id", middleware.Authentication(), controllers.GetTable())
	incomingRoutes.PATCH("/table/:table_id", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.UpdateTable())
}