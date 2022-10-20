package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		// if auth not bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := NewResponseAPI("Unauthorized", "error", http.StatusUnauthorized, nil)
			ctx.JSON(
				http.StatusUnauthorized,
				response,
			)
			return
		}
		

	}
}
