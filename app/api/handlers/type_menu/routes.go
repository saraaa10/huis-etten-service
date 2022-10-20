package typemenu

import (
	"service-api/app/core/type_menu/services"

	"github.com/gin-gonic/gin"
)

func NewTypeMenuRoutes(router *gin.RouterGroup, service services.TypeMenuService) {
	handler := NewTypeMenuHandler(service)
	typemenu := router.Group("/type-menu")
	{
		typemenu.GET("/", handler.GetAllTypeMenu)
	}
}
