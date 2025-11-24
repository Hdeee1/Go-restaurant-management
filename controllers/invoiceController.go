package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetInvoices() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invoices []models.Invoice

		if err := database.DB.Find(&invoices).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"invoices": invoices,
		})
	}
}

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

func CreateInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invoice models.Invoice

		invoice.Invoice_id = uuid.New().String()

		if err := ctx.BindJSON(&invoice); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Create(&invoice).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "invoice created",
			"invoice_id": invoice.Invoice_id,
		})
	}
}

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
			"message": "invoice updated",
			"invoice_id": invoice.Invoice_id,
		})
	}
}
