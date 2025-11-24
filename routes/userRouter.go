package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hdeee1/go-restaurant-management/controllers"
	)

func UserRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}