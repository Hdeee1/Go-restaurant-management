package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hdeee1/go-restaurant-management/controllers"
	"github.com/Hdeee1/go-restaurant-management/middleware"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/invoices", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.CreateInvoice())
	incomingRoutes.GET("/invoices", middleware.Authentication(), controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", middleware.Authentication(), controllers.GetInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", middleware.Authentication(), middleware.CheckRole("ADMIN"), controllers.UpdateInvoice())
}