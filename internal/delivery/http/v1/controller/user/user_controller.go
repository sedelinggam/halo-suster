package userControllers

import (
	userHandler "halo-suster/internal/delivery/http/v1/controller/user/handler"
	userService "halo-suster/internal/service/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, userSvc userService.UserService) {
	user := group.Group("/user")
	handler := userHandler.NewHandler(userSvc, val)

	publicRoute := user
	publicRoute.Use()
	publicRoute.POST("/register", handler.Register)
	publicRoute.POST("/login", handler.Login)
}
