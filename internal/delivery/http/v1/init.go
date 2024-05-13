package v1

import (
	staffControllers "halo-suster/internal/delivery/http/v1/controller/staff"
	staffService "halo-suster/internal/service/staff"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "halo-suster/cmd/docs" // docs is generated by Swag CLI, you have to import it.
)

func Init(app *echo.Echo, db *sqlx.DB, val *validator.Validate) {
	var (
		staffSvc = staffService.New(db)
	)
	v1 := app.Group("/v1")
	staffControllers.Init(v1, val, staffSvc)
	v1.GET("/swagger/*", echoSwagger.WrapHandler)
}
