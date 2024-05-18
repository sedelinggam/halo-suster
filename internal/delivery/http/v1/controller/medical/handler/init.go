package medicalHandler

import (
	medicalService "halo-suster/internal/service/medical"

	"github.com/go-playground/validator/v10"
)

type medicalHandler struct {
	medicalService medicalService.MedicalService
	val            *validator.Validate
}

func NewHandler(medicalService medicalService.MedicalService, val *validator.Validate) *medicalHandler {
	return &medicalHandler{
		medicalService: medicalService,
		val:            val,
	}
}
