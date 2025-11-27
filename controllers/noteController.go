package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
)

// GetNotes godoc
// @Summary Get all notes
// @Description Retrieve a paginated list of all notes
// @Tags Notes
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notes [get]
func GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var notes []models.Note

		result := database.DB.Scopes(helpers.Paginate(ctx)).Find(&notes)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"notes": notes,
			"page":  ctx.DefaultQuery("page", "1"),
			"limit": ctx.DefaultQuery("limit", "10"),
		})
	}
}

// GetNote godoc
// @Summary Get note by ID
// @Description Retrieve a specific note by note_id
// @Tags Notes
// @Accept json
// @Produce json
// @Param note_id path string true "Note ID"
// @Security BearerAuth
// @Success 200 {object} models.Note
// @Failure 404 {object} map[string]interface{}
// @Router /notes/{note_id} [get]
func GetNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		note_id := ctx.Param("note_id")

		var note models.Note

		if err := database.DB.Where("note_id = ?", note_id).First(&note).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "note_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, note)
	}
}

// CreateNote godoc
// @Summary Create a new note
// @Description Create a new note with the provided information
// @Tags Notes
// @Accept json
// @Produce json
// @Param note body models.Note true "Note object"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notes [post]
func CreateNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var note models.Note

		if err := ctx.BindJSON(&note); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Create(&note).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "note created",
			"note_id": note.Note_id,
		})
	}
}

// UpdateNote godoc
// @Summary Update a note
// @Description Update an existing note by note_id
// @Tags Notes
// @Accept json
// @Produce json
// @Param note_id path string true "Note ID"
// @Param note body models.Note true "Note object"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notes/{note_id} [put]
func UpdateNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		note_id := ctx.Param("note_id")

		var note models.Note

		if err := database.DB.Where("note_id = ?", note_id).First(&note).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "note_id not found"})
			return
		}

		var updateData models.Note
		if err := ctx.BindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Model(&note).Updates(updateData).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "note updated",
			"note_id": note.Note_id,
		})
	}
}
