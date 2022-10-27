package util

import (
	"net/http"
	"service-api/app/api/auth"
	"service-api/app/core/users/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(serviceAuth auth.Service, userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		// if auth not bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := NewResponseAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response,
			)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := serviceAuth.ValidateToken(tokenString)
		if err != nil {
			response := NewResponseAPI("Unauthorized. can't validate token", "error", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response,
			)
			return
		}

		// check token
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := NewResponseAPI("Unauthorized. Token not valid", "error", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response,
			)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(uint(userID))
		if err != nil {
			response := NewResponseAPI("Unauthorized. User not found", "error", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response,
			)
			return
		}

		ctx.Set("currentUser", user)
		ctx.Set("token", tokenString)
	}
}
