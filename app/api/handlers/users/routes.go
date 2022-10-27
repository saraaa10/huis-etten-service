package users

import (
	"service-api/app/api/auth"
	"service-api/app/core/users/services"

	"github.com/gin-gonic/gin"
)

func NewUserRoute(routes *gin.RouterGroup, services services.UserService, authService auth.Service) {
	handler := NewUserHandler(services, authService)
	user := routes.Group("/users")
	{
		user.POST("/register", handler.RegisterUser)
		user.GET("/", handler.GetAllUsers)
		user.POST("/login", handler.LoginUser)
	}
}
