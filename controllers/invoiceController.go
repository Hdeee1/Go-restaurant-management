package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetInvoices godoc
// @Summary Get all invoices
// @Description Retrieve a paginated list of all invoices
// @Tags Invoices
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invoices [get]
func GetInvoices() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invoices []models.Invoice

		result := database.DB.Scopes(helpers.Paginate(ctx)).Find(&invoices)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"invoices": invoices,
			"page":     ctx.DefaultQuery("page", "1"),
			"limit":    ctx.DefaultQuery("limit", "10"),
		})
	}
}

// GetInvoice godoc
// @Summary Get invoice by ID
// @Description Retrieve a specific invoice by invoice_id
// @Tags Invoices
// @Accept json
// @Produce json
// @Param invoice_id path string true "Invoice ID"
// @Security BearerAuth
// @Success 200 {object} models.Invoice
// @Failure 404 {object} map[string]interface{}
// @Router /invoices/{invoice_id} [get]
func GetInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoice_id := ctx.Param("invoice_id")

		var invoice models.Invoice

		if err := database.DB.Where("invoice_id = ?", invoice_id).First(&invoice).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invoice_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, invoice)
	}
}

// CreateInvoice godoc
// @Summary Create a new invoice
// @Description Create a new invoice with the provided information
// @Tags Invoices
// @Accept json
// @Produce json
// @Param invoice body models.Invoice true "Invoice object"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invoices [post]
func CreateInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invoice models.Invoice

		if err := ctx.BindJSON(&invoice); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := helpers.Validate.Struct(invoice); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		invoice.Invoice_id = uuid.New().String()

		if err := database.DB.Create(&invoice).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message":    "invoice created",
			"invoice_id": invoice.Invoice_id,
		})
	}
}

// UpdateInvoice godoc
// @Summary Update an invoice
// @Description Update an existing invoice by invoice_id
// @Tags Invoices
// @Accept json
// @Produce json
// @Param invoice_id path string true "Invoice ID"
// @Param invoice body models.Invoice true "Invoice object"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invoices/{invoice_id} [put]
func UpdateInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoice_id := ctx.Param("invoice_id")

		var invoice models.Invoice

		if err := database.DB.Where("invoice_id = ?", invoice_id).First(&invoice).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invoice_id not found"})
			return
		}

		var updateData models.Invoice
		if err := ctx.BindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Model(&invoice).Updates(updateData).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":    "invoice updated",
			"invoice_id": invoice.Invoice_id,
		})
	}
}
