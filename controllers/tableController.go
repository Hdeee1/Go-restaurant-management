package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetTables godoc
// @Summary Get all tables
// @Description Retrieve a paginated list of all tables
// @Tags Tables
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tables [get]
func GetTables() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tables []models.Table

		result := database.DB.Scopes(helpers.Paginate(ctx)).Find(&tables)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"tables": tables,
			"page":   ctx.DefaultQuery("page", "1"),
			"limit":  ctx.DefaultQuery("limit", "10"),
		})
	}
}

// GetTable godoc
// @Summary Get table by ID
// @Description Retrieve a specific table by table_id
// @Tags Tables
// @Accept json
// @Produce json
// @Param table_id path string true "Table ID"
// @Security BearerAuth
// @Success 200 {object} models.Table
// @Failure 404 {object} map[string]interface{}
// @Router /tables/{table_id} [get]
func GetTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tableID := ctx.Param("table_id")

		var table models.Table

		if err := database.DB.Where("table_id = ?", tableID).First(&table).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "table_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, table)
	}
}

// CreateTable godoc
// @Summary Create a new table
// @Description Create a new table with the provided information
// @Tags Tables
// @Accept json
// @Produce json
// @Param table body models.Table true "Table object"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tables [post]
func CreateTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var table models.Table

		if err := ctx.BindJSON(&table); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := helpers.Validate.Struct(table); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		table.Table_id = uuid.New().String()

		if err := database.DB.Create(&table).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message":  "table created",
			"table_id": table.Table_id,
		})
	}
}

// UpdateTable godoc
// @Summary Update a table
// @Description Update an existing table by table_id
// @Tags Tables
// @Accept json
// @Produce json
// @Param table_id path string true "Table ID"
// @Param table body models.Table true "Table object"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tables/{table_id} [put]
func UpdateTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tableID := ctx.Param("table_id")

		var table models.Table

		if err := database.DB.Where("table_id = ?", tableID).First(&table).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "table_id not found"})
			return
		}

		var updateData models.Table
		if err := ctx.BindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Model(&table).Updates(updateData).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "table updated",
			"table_id": table.Table_id,
		})
	}
}
