package typemenu

import (
	"net/http"
	"service-api/app/core/type_menu/services"
	"service-api/common/util"

	"github.com/gin-gonic/gin"
)

type typeMenuHandler struct {
	service services.TypeMenuService
}

func NewTypeMenuHandler(s services.TypeMenuService) *typeMenuHandler {
	return &typeMenuHandler{
		service: s,
	}
}

func (handler *typeMenuHandler) GetAllTypeMenu(ctx *gin.Context) {
	typeMenus, err := handler.service.FindAll()
	if err != nil {
		response := util.NewResponseAPI(err.Error(), "error", http.StatusBadRequest, nil)
		ctx.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	if len(typeMenus) == 0 {
		response := util.NewResponseAPIList("Type Menu Not Found", "error", http.StatusNotFound, nil, 0)
		ctx.JSON(
			http.StatusNotFound,
			response,
		)

		return
	}

	response := util.NewResponseAPIList("Success Fetch Type Menu", "success", http.StatusOK, typeMenus, len(typeMenus))
	ctx.JSON(
		http.StatusOK,
		response,
	)
}
