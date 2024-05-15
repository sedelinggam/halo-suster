package userControllers

import (
	userHandler "halo-suster/internal/delivery/http/v1/controller/user/handler"
	userService "halo-suster/internal/service/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, userSvc userService.UserService, jwt echo.MiddlewareFunc) {
	user := group.Group("/user")
	handler := userHandler.NewHandler(userSvc, val)

	//IT
	itRouterPublic := user.Group("/it")
	itRouterPublic.POST("/register", handler.ITRegister)
	itRouterPublic.POST("/login", handler.ITLogin)

	//Nurse
	nursePublic := user.Group("/nurse")
	nursePublic.POST("/login", handler.NurseLogin)

	//Private
	itRouterPrivate := user
	itRouterPrivate.Use(jwt)
	itRouterPrivate.POST("/nurse/register", handler.NurseRegister)
	itRouterPrivate.GET("", handler.User)
	itRouterPrivate.PUT("/nurse/:userId", handler.NurseUpdate)
	itRouterPrivate.DELETE("/nurse/:userId", handler.NurseDelete)
	itRouterPrivate.POST("/nurse/:userId/access", handler.NursePassword)
}
