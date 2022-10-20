package categorymenu

import (
	"net/http"
	"service-api/app/core/category_menu/services"
	"service-api/common/util"

	"github.com/gin-gonic/gin"
)

type categoryMenuHandler struct {
	service services.CategoryMenuService
}

func NewcategoryMenuHandler(s services.CategoryMenuService) *categoryMenuHandler {
	return &categoryMenuHandler{
		service: s,
	}
}

func (handler *categoryMenuHandler) GetAllCategoryMenu(context *gin.Context) {
	categoryMenus, err := handler.service.FindAll()
	if err != nil {
		errMsg := gin.H{
			"error": err.Error(),
		}

		response := util.NewResponseAPI("Failed Fetch Data Category Menu", "error", http.StatusBadRequest, errMsg)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	if len(categoryMenus) == 0 {
		errMsg := gin.H{
			"error": "Data Category Menu Not Found",
		}
		response := util.NewResponseAPI("Failed Fetch Data Category Menu", "error", http.StatusNotFound, errMsg)
		context.JSON(http.StatusNotFound, response)
		return
	}

	response := util.NewResponseAPIList("Success Fetch Data Category Menu", "success", http.StatusOK, categoryMenus, len(categoryMenus))
	context.JSON(http.StatusOK, response)
}
