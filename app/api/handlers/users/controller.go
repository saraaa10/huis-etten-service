package users

import (
	"net/http"
	"service-api/app/api/auth"
	"service-api/app/core/users/services"
	"service-api/app/util/users"
	"service-api/common/util"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     services.UserService
	authService auth.Service
}

func NewUserHandler(service services.UserService, auth auth.Service) *userHandler {
	return &userHandler{
		service:     service,
		authService: auth,
	}
}

func (handler *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := util.NewResponseAPI(errMsg, "error", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := handler.service.RegisterUser(input)
	if err != nil {
		errMsg := err.Error()
		response := util.NewResponseAPI(errMsg, "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	

	token, err := handler.authService.GenerateToken(int(newUser.ID))
	if err != nil {
		errMsg := err.Error()
		response := util.NewResponseAPI(errMsg, "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(newUser, token)
	response := util.NewResponseAPI("success register user", "success", http.StatusCreated, formatter)
	c.JSON(http.StatusCreated, response)
}

func (handler *userHandler) LoginUser(c *gin.Context) {
	var input users.LoginUserInput

	// binding input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := util.NewResponseAPI(errMsg, "error", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// process login
	userLogin, err := handler.service.LoginUser(input)
	if err != nil {
		errMsg := err.Error()
		response := util.NewResponseAPI(errMsg, "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// set token
	token, err := handler.authService.GenerateToken(int(userLogin.ID))
	if err != nil {
		errMsg := err.Error()
		response := util.NewResponseAPI(errMsg, "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	format := users.FormatUser(userLogin, token)
	response := util.NewResponseAPI("success login user", "success", http.StatusOK, format)
	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) GetAllUsers(c *gin.Context) {
	users, err := handler.service.FindAll()
	if err != nil {
		errMsg := err.Error()
		response := util.NewResponseAPI(errMsg, "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.NewResponseAPI("success get all users", "success", http.StatusOK, users)
	c.JSON(http.StatusOK, response)
}
