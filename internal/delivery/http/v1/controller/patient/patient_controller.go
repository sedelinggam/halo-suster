package patientController

import (
	patientHandler "halo-suster/internal/delivery/http/v1/controller/patient/handler"
	patientService "halo-suster/internal/service/patient"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, patientSvc patientService.PatientService, jwt echo.MiddlewareFunc) {
	patient := group.Group("/medical/patient")
	handler := patientHandler.NewHandler(patientSvc, val)

	//Private
	itRouterPrivate := patient
	itRouterPrivate.Use(jwt)
	itRouterPrivate.POST("", handler.CreatePatient)
	itRouterPrivate.GET("", handler.CreatePatient)
}
