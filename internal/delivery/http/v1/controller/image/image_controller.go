package imageController

import (
	imageHandler "halo-suster/internal/delivery/http/v1/controller/image/handler"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, jwt echo.MiddlewareFunc) {
	image := group.Group("/image")
	handler := imageHandler.NewHandler(val)

	privateRoute := image
	// TODO: Add middleware
	privateRoute.Use(jwt)
	privateRoute.POST("", handler.StoreImage)
}
