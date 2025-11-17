package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Hdeee1/go-restaurant-management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incomingRoutes.GET("/invoices", controllers.CreateInvoice())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.UpdateInvoice())
}