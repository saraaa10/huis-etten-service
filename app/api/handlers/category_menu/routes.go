package categorymenu

import (
	"service-api/app/core/category_menu/services"

	"github.com/gin-gonic/gin"
)

func NewCategoryMenuRoutes(router *gin.RouterGroup, service services.CategoryMenuService) {
	handler := NewcategoryMenuHandler(service)

	categoryMenuRoutes := router.Group("/category-menu")
	{
		categoryMenuRoutes.GET("/", handler.GetAllCategoryMenu)
	}
}
