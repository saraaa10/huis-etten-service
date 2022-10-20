package main

import (
	"fmt"
	categorymenu "service-api/app/api/handlers/category_menu"
	typemenu "service-api/app/api/handlers/type_menu"
	usertype "service-api/app/api/handlers/user_type"

	userTypeRepoImpl "service-api/app/core/user_type/repositories/impl"
	userTypeServiceImpl "service-api/app/core/user_type/services/impl"

	categoryMenuRepoImpl "service-api/app/core/category_menu/repositories/impl"
	categoryMenuServiceImpl "service-api/app/core/category_menu/services/impl"

	typeMenuRepoImpl "service-api/app/core/type_menu/repositories/impl"
	typeMenuServiceImpl "service-api/app/core/type_menu/services/impl"

	"service-api/common/constant"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	db := constant.InitDB()
	constant.AutoMigrate(db)

	// User Type Repo & Service
	userTypeRepo := userTypeRepoImpl.NewUserTypeRepository(db)
	userTypeService := userTypeServiceImpl.NewUserTypeService(userTypeRepo)

	// category Menu Repo & Service
	categoryMenuRepo := categoryMenuRepoImpl.NewCategoryMenuRepository(db)
	categoryMenuService := categoryMenuServiceImpl.NewCategoryMenuService(categoryMenuRepo)

	// type menu Repo & Service
	typeMenuRepo := typeMenuRepoImpl.NewTypeMenuRepository(db)
	typeMenuService := typeMenuServiceImpl.NewTypeMenuService(typeMenuRepo)

	// router api/V1 Configuration
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		usertype.NewUserTypeRoutes(api, userTypeService)
		categorymenu.NewCategoryMenuRoutes(api, categoryMenuService)
		typemenu.NewTypeMenuRoutes(api, typeMenuService)
	}

	router.Run(":8080")
}
