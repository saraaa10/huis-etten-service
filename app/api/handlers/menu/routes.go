package menu

import (
	"service-api/app/core/menu/services"

	"github.com/gin-gonic/gin"
)

func NewMenuRoutes(router *gin.RouterGroup, service services.MenuService) {
	handler := NewMenuHandler(service)

	menu := router.Group("/menu")
	{
		menu.GET("/", handler.GetAllMenu)
	}
}