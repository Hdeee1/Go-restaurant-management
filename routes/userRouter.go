package routes

import (
	"github.com/Hdeee1/go-restaurant-management/controllers"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.GET("/users", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", middleware.Authentication(), controllers.GetUser())
}