package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
)

func GetNotes() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		var notes []models.Note

		result := database.DB.Scopes(helpers.Paginate(ctx)).Find(&notes)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return 
		}

		ctx.JSON(http.StatusOK, gin.H{
			"notes": notes,
			"page": ctx.DefaultQuery("page", "1"),
			"limit": ctx.DefaultQuery("limit", "10"),
		})
	}
}

func GetNote() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		note_id := ctx.Param("note_id")

		var note models.Note

		if err := database.DB.Where("note_id = ?", note_id).First(&note).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "note_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, note)	
	}
}

func CreateNote() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
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

func UpdateNote() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
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