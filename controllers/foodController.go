package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetFoods godoc
//
//	@Summary		Get all foods
//	@Description	Get all foods
//	@Tags			Foods
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	false	"Page number"
//	@Param			limit	query		int	false	"Limit"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Failure		401		{object}	map[string]interface{}
//	@Security		BearerAuth
//	@Router			/foods [get]
func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var foods []models.Food

		result := database.DB.Scopes(helpers.Paginate(ctx)).Find(&foods)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"foods": foods,
			"page":  ctx.DefaultQuery("page", "1"),
			"limit": ctx.DefaultQuery("limit", "10"),
		})
	}
}

// GetFood godoc
//
//	@Summary		Get a food by ID
//	@Description	Get a food by ID
//	@Tags			Foods
//	@Accept			json
//	@Produce		json
//	@Param			food_id	path		string	true	"Food ID"
//	@Success		200		{object}	models.Food
//	@Failure		404		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Router			/foods/{food_id} [get]
func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		foodID := ctx.Param("food_id")

		var food models.Food

		if err := database.DB.Where("food_id = ?", foodID).First(&food).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "food_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, food)
	}
}

// AddFood godoc
//
//	@Summary		Add a new food
//	@Description	Add a new food
//	@Tags			Foods
//	@Accept			json
//	@Produce		json
//	@Param			food	body		models.Food	true	"Food object"
//	@Success		201		{object}	models.Food
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Router			/foods [post]
func AddFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var food models.Food

		if err := ctx.BindJSON(&food); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := helpers.Validate.Struct(food); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		food.Food_id = uuid.New().String()

		var menu models.Menu
		err := database.DB.Where("menu_id = ?", *food.Menu_id).First(&menu).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "menu_id not found"})
			return
		}

		if err := database.DB.Create(&food).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Food created",
			"food_id": food.Food_id,
		})
	}
}

// UpdateFood godoc
//
//	@Summary		Update a food
//	@Description	Update a food
//	@Tags			Foods
//	@Accept			json
//	@Produce		json
//	@Param			food_id	path		string		true	"Food ID"
//	@Param			food	body		models.Food	true	"Food object"
//	@Success		200		{object}	models.Food
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Router			/foods/{food_id} [put]
func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		foodID := ctx.Param("food_id")

		var food models.Food

		if err := database.DB.Where("food_id = ?", foodID).First(&food).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "food_id not found"})
			return
		}

		var updateData models.Food
		if err := ctx.BindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Model(&food).Updates(updateData).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Food updated",
			"food":    food,
		})
	}
}
