package usertype

import (
	"net/http"
	"service-api/app/core/user_type/services"
	"service-api/common/util"

	"github.com/gin-gonic/gin"
)

type userTypeHandler struct {
	service services.UserTypeService
}

func NewUserTypeHandler(service services.UserTypeService) *userTypeHandler {
	return &userTypeHandler{
		service: service,
	}
}

func (handler *userTypeHandler) GetAllUserType(c *gin.Context) {
	userType, err := handler.service.GetAll()
	if err != nil {
		errMsg := gin.H{
			"error": err.Error(),
		}

		response := util.NewResponseAPI("Failed Fetch User Type", "error", http.StatusBadRequest, errMsg)
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	if len(userType) == 0 {
		response := util.NewResponseAPIList("User Type Not Found", "error", http.StatusNotFound, userType, 0)
		c.JSON(
			http.StatusNotFound,
			response,
		)
		return
	}

	response := util.NewResponseAPIList("Success Fetch User Type", "success", http.StatusOK, userType, len(userType))
	c.JSON(
		http.StatusOK,
		response,
	)
}
