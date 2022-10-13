package usertype

import (
	"service-api/app/core/user_type/services"

	"github.com/gin-gonic/gin"
)

func NewUserTypeRoutes(router *gin.RouterGroup, service services.UserTypeService) {
	userTypeHandler := NewUserTypeHandler(service)
	userType := router.Group("/user-type")
	{
		userType.GET("/", userTypeHandler.GetAllUserType)
	}
}
