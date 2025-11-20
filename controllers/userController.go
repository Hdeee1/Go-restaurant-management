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

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []models.User

		err := database.DB.Find(&users).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		for i := range users {
			users[i].Password = nil
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": users,
		})
	}
}

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

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		
		user.User_id = uuid.New().String()

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		hashedPass := HashPassword(*user.Password)
		user.Password = &hashedPass

		token, refreshToken, err := helpers.GenerateToken(*user.Email, user.User_id)
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

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginInput struct {
			Email		string	`json:"email"`
			Password	string	`json:"password"`
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

		token, refreshToken, err := helpers.GenerateToken(loginInput.Email, user.User_id)
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
			"token": token,
			"user_id": user.User_id, 
		})
	}
}