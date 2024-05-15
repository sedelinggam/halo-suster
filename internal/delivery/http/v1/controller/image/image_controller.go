package imageController

import (
	imageHandler "halo-suster/internal/delivery/http/v1/controller/image/handler"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate) {
	image := group.Group("/image")
	handler := imageHandler.NewHandler(val)

	publicRoute := image
	// TODO: Add middleware
	publicRoute.POST("", handler.StoreImage)
}
