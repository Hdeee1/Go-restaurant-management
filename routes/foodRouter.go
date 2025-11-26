package routes

import (
	"github.com/Hdeee1/go-restaurant-management/controllers"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/foods", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.AddFood())
	incomingRoutes.PATCH("/foods/:food_id", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.UpdateFood())
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFood())
}