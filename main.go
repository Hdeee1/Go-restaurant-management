package main

import (
	"os"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/Hdeee1/go-restaurant-management/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	database.InitDB()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRouter(router)
	router.Use(middleware.Authentication())
	
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)
}