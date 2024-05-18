package medicalControllers

import (
	medicalHandler "halo-suster/internal/delivery/http/v1/controller/medical/handler"
	medicalService "halo-suster/internal/service/medical"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, medicalSvc medicalService.MedicalService, jwt echo.MiddlewareFunc) {
	medical := group.Group("/medical/record")
	handler := medicalHandler.NewHandler(medicalSvc, val)

	//Private
	itRouterPrivate := medical
	itRouterPrivate.Use(jwt)
	itRouterPrivate.POST("", handler.CreateMedicalRecord)
	itRouterPrivate.GET("", handler.GetMedicalRecords)
}
