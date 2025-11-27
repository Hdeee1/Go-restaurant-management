package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetMenus godoc
//
//	@Summary		Get all menus
//	@Description	Retrieve a paginated list of all menus
//	@Tags			Menus
//	@Accept			json
//	@Produce		json
//	@Param			page	query	int	false	"Page number"		default(1)
//	@Param			limit	query	int	false	"Items per page"	default(10)
//	@Security		BearerAuth
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/menus [get]
func GetMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menus []models.Menu

		result := database.DB.Scopes(helpers.Paginate(ctx)).Find(&menus)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"menus": menus,
			"page":  ctx.DefaultQuery("page", "1"),
			"limit": ctx.DefaultQuery("limit", "10"),
		})
	}
}

// GetMenu godoc
//
//	@Summary		Get menu by ID
//	@Description	Retrieve a specific menu by menu_id
//	@Tags			Menus
//	@Accept			json
//	@Produce		json
//	@Param			menu_id	path	string	true	"Menu ID"
//	@Security		BearerAuth
//	@Success		200	{object}	models.Menu
//	@Failure		404	{object}	map[string]interface{}
//	@Router			/menus/{menu_id} [get]
func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		menu_id := ctx.Param("menu_id")

		var menu models.Menu

		if err := database.DB.Where("menu_id = ?", menu_id).First(&menu).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "menu_id not found"})
			return
		}

		ctx.JSON(http.StatusOK, menu)
	}
}

// CreateMenu godoc
//
//	@Summary		Create a new menu (Admin only)
//	@Description	Create a new menu with the provided information. Requires admin role.
//	@Tags			Menus
//	@Accept			json
//	@Produce		json
//	@Param			menu	body	models.Menu	true	"Menu object"
//	@Security		BearerAuth
//	@Success		201	{object}	map[string]interface{}
//	@Failure		400	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/menus [post]
func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menu models.Menu

		if err := ctx.BindJSON(&menu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := helpers.Validate.Struct(menu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menu.Menu_id = uuid.New().String()

		if menu.Start_date != nil && menu.End_date != nil {
			if menu.End_date.Before(*menu.Start_date) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "end_date must be after start_date"})
				return
			}
		}

		if err := database.DB.Create(&menu).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "menu created",
			"menu_id": menu.Menu_id,
		})
	}
}

// UpdateMenu godoc
//
//	@Summary		Update a menu (Admin only)
//	@Description	Update an existing menu by menu_id. Requires admin role.
//	@Tags			Menus
//	@Accept			json
//	@Produce		json
//	@Param			menu_id	path	string		true	"Menu ID"
//	@Param			menu	body	models.Menu	true	"Menu object"
//	@Security		BearerAuth
//	@Success		200	{object}	map[string]interface{}
//	@Failure		400	{object}	map[string]interface{}
//	@Failure		404	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/menus/{menu_id} [put]
func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		menuID := ctx.Param("menu_id")

		var menu models.Menu

		if err := database.DB.Where("menu_id = ?", menuID).First(&menu).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "menu_id not found"})
			return
		}

		var updateData models.Menu
		if err := ctx.BindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Model(&menu).Updates(updateData).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "menu updated",
			"menu_id": menu.Menu_id,
		})
	}
}
