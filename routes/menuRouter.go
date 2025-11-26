package routes

import (
	"github.com/Hdeee1/go-restaurant-management/controllers"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/menus", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.CreateMenu())
	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	incomingRoutes.PATCH("/menus/:menu_id", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.UpdateMenu())
}
