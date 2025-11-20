package middleware

import (
	"net/http"
	"strings"

	"github.com/Hdeee1/go-restaurant-management/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
			return 
		}

		tokenString := strings.TrimPrefix(token, "Bearer")

		claims, err := helpers.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return 
		}

		ctx.Set("email", claims.Email)
		ctx.Set("user_id", claims.User_id)

		ctx.Next()
	}
}