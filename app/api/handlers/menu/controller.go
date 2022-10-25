package menu

import (
	"net/http"
	"service-api/app/core/menu/services"
	"service-api/common/util"

	"github.com/gin-gonic/gin"
)

type menuHandler struct {
	service services.MenuService
}

func NewMenuHandler(service services.MenuService) *menuHandler {
	return &menuHandler{service: service}
}

func (handler *menuHandler) GetAllMenu(ctx *gin.Context) {
	menu, err := handler.service.FindAll()
	if err != nil {
		errMsg := err.Error()
		response := util.NewResponseAPI(errMsg, "error", http.StatusBadRequest, nil)
		ctx.JSON(
			http.StatusBadRequest,
			response,
		)

		return
	}

	// if menu null
	if len(menu) == 0 {
		response := util.NewResponseAPIList("Menu not found", "error", http.StatusNotFound, nil, 0)
		ctx.JSON(
			http.StatusOK,
			response,
		)

		return
	}

	response := util.NewResponseAPIList("Success get all menu", "success", http.StatusOK, menu, len(menu))
	ctx.JSON(
		http.StatusOK,
		response,
	)
}