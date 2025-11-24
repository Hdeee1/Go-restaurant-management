package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hdeee1/go-restaurant-management/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/foods", controllers.AddFood())
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFood())
	incomingRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
}