package main

import (
	"fmt"

	// routes & handler
	auth "service-api/app/api/auth"
	categorymenu "service-api/app/api/handlers/category_menu"
	menu "service-api/app/api/handlers/menu"
	typemenu "service-api/app/api/handlers/type_menu"
	usertype "service-api/app/api/handlers/user_type"
	user "service-api/app/api/handlers/users"

	// repo & service
	categoryMenuRepoImpl "service-api/app/core/category_menu/repositories/impl"
	categoryMenuServiceImpl "service-api/app/core/category_menu/services/impl"
	menuRepoImpl "service-api/app/core/menu/repositories/impl"
	menuServiceImpl "service-api/app/core/menu/services/impl"
	typeMenuRepoImpl "service-api/app/core/type_menu/repositories/impl"
	typeMenuServiceImpl "service-api/app/core/type_menu/services/impl"
	userTypeRepoImpl "service-api/app/core/user_type/repositories/impl"
	userTypeServiceImpl "service-api/app/core/user_type/services/impl"
	userRepoImpl "service-api/app/core/users/repositories/impl"
	userServiceImpl "service-api/app/core/users/services/impl"

	// others
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

	// menu Repo & Service
	menuRepo := menuRepoImpl.NewMenuRepository(db)
	menuService := menuServiceImpl.NewMenuService(menuRepo)

	// auth service
	authService := auth.NewService()

	// user repo & service
	userRepo := userRepoImpl.NewUserRepository(db)
	userService := userServiceImpl.NewUserService(userRepo, userTypeRepo)

	// router api/V1 Configuration
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		user.NewUserRoute(api, userService, authService)
		usertype.NewUserTypeRoutes(api, userTypeService)
		categorymenu.NewCategoryMenuRoutes(api, categoryMenuService)
		typemenu.NewTypeMenuRoutes(api, typeMenuService)
		menu.NewMenuRoutes(api, menuService)
	}

	// ping request
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to service-api",
		})
	})

	// Run Server
	router.Run(":3000")
}
