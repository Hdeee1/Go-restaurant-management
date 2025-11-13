package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hdeee1/go-restaurant-management/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFood())
	incomingRoutes.POST("/foods", controllers.AddFood())
	incomingRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
}