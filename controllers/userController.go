package controllers

import (
	"net/http"

	"github.com/Hdeee1/go-restaurant-management/database"
	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/Hdeee1/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func VerifyPassword(password string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return false, "Wrong password"
	}

	return true, ""
}

// GetUsers godoc
//
//	@Summary		Get all users (Admin only)
//	@Description	Retrieve a paginated list of all users
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			page	query	int	false	"Page number"		default(1)
//	@Param			limit	query	int	false	"Items per page"	default(10)
//	@Security		BearerAuth
//	@Success		200	{object}	models.UsersListResponse
//	@Failure		401	{object}	models.ErrorResponse
//	@Failure		500	{object}	models.ErrorResponse
//	@Router			/users [get]
func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []models.User

		result := database.DB.Scopes(helpers.Paginate(ctx)).Find(&users)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":  users,
			"page":  ctx.DefaultQuery("page", "1"),
			"limit": ctx.DefaultQuery("limit", "10"),
		})
	}
}

// GetUser godoc
//
//	@Summary		Get user by ID
//	@Description	Retrieve a specific user by their user_id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	string	true	"User ID"
//	@Security		BearerAuth
//	@Success		200	{object}	models.UserResponse
//	@Failure		404	{object}	models.ErrorResponse
//	@Router			/users/{user_id} [get]
func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")

		var user models.User

		if err := database.DB.Where("user_id = ?", userID).Find(&user).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		user.Password = nil

		ctx.JSON(http.StatusOK, user)
	}
}

// SignUp godoc
//
//	@Summary		Register a new user
//	@Description	Create a new user account with email, password, and personal information
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.SignUpRequest	true	"User registration details"
//	@Success		201		{object}	models.SignUpResponse
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/users/signup [post]
func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := helpers.Validate.Struct(user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.User_id = uuid.New().String()
		hashedPass := HashPassword(*user.Password)
		user.Password = &hashedPass

		token, refreshToken, err := helpers.GenerateToken(*user.Email, user.User_id, *user.Role)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token"})
			return
		}

		user.Token = &token
		user.Refresh_Token = &refreshToken

		if err := database.DB.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Successfully signed up",
			"user_id": user.User_id,
		})
	}
}

// Login godoc
//
//	@Summary		User login
//	@Description	Authenticate user with email and password, returns JWT token
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		models.LoginRequest	true	"Login credentials"
//	@Success		200			{object}	models.LoginResponse
//	@Failure		400			{object}	models.ErrorResponse
//	@Failure		401			{object}	models.ErrorResponse
//	@Router			/users/login [post]
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginInput struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := ctx.BindJSON(&loginInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User

		err := database.DB.Where("email = ?", loginInput.Email).First(&user).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		isValid, msg := VerifyPassword(*user.Password, loginInput.Password)
		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		token, refreshToken, err := helpers.GenerateToken(loginInput.Email, user.User_id, *user.Role)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		user.Token = &token
		user.Refresh_Token = &refreshToken

		err = database.DB.Save(&user).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
			"token":   token,
			"user_id": user.User_id,
		})
	}
}
