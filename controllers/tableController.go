package controllers

import (
	"net/http"
	
	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTables() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		var tables []models.Table

		if err := database.DB.Find(&tables).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"tables": tables,
		})
	}
}

func GetTable() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		
	}
}

func CreateTable() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		var table models.Table

		table.Table_id = uuid.New().String()

		if err := ctx.BindJSON(&table); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Create(&table).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "table created",
			"table_id": table.Table_id,
		})
	}
}

func UpdateTable() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		
	}
}