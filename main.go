package main

import (
	"fmt"
	usertype "service-api/app/api/handlers/user_type"
	userTypeRepoImpl "service-api/app/core/user_type/repositories/impl"
	userTypeServiceImpl "service-api/app/core/user_type/services/impl"
	"service-api/common/constant"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	db := constant.InitDB()

	userTypeRepo := userTypeRepoImpl.NewUserTypeRepository(db)
	userTypeService := userTypeServiceImpl.NewUserTypeService(userTypeRepo)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		usertype.NewUserTypeRoutes(api, userTypeService)
	}

	router.Run(":8080")
}
