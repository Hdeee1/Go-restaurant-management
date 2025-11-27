package main

import (
	"fmt"
	"os"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/middleware"
	"github.com/Hdeee1/go-restaurant-management/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Hdeee1/go-restaurant-management/docs"
)

// @title Restaurant Management API
// @version 1.0
// @description This is a comprehensive restaurant management system API built with Go and Gin framework
// @description Features include user management, food menus, table reservations, orders, and invoicing

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8081
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func printRoutes(router *gin.Engine) {
	routesList := router.Routes()

	// Group routes by category
	routeGroups := make(map[string][]gin.RouteInfo)
	for _, route := range routesList {
		// Determine category based on path prefix
		category := "other"
		if len(route.Path) > 1 {
			path := route.Path[1:] // Remove leading /
			for idx := 0; idx < len(path); idx++ {
				if path[idx] == '/' {
					category = path[:idx]
					break
				}
			}
			if category == "other" {
				category = path
			}
		}
		routeGroups[category] = append(routeGroups[category], route)
	}

	// Define the order of categories based on application flow
	categoryOrder := []string{"users", "food", "menu", "table", "order", "orderitem", "invoice"}

	// Print routes in order
	for _, category := range categoryOrder {
		if routes, exists := routeGroups[category]; exists {
			fmt.Printf("\n%s ROUTES\n", toUpperCategory(category))
			for _, route := range routes {
				fmt.Printf("  %-8s %s\n", route.Method, route.Path)
			}
			fmt.Println("─────────────────────────────────────")
		}
	}

	// Print any other routes not in the defined order
	for category, routes := range routeGroups {
		if !contains(categoryOrder, category) {
			fmt.Printf("\n%s ROUTES\n", toUpperCategory(category))
			for _, route := range routes {
				fmt.Printf("  %-8s %s\n", route.Method, route.Path)
			}
			fmt.Println("─────────────────────────────────────")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("\nServer running on http://localhost:%s\n", port)
	fmt.Println("════════════════════════════════════════\n")
}

func toUpperCategory(s string) string {
	if len(s) == 0 {
		return s
	}
	// Convert first letter to uppercase
	result := []rune(s)
	if result[0] >= 'a' && result[0] <= 'z' {
		result[0] = result[0] - 32
	}
	return string(result)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	database.InitDB()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Swagger documentation route (accessible without authentication)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRouter(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	// Print all registered routes
	printRoutes(router)

	router.Run(":" + port)
}
