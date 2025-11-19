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

	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
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

	}
}